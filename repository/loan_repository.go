package repository

import (
	"billing/config"
	"billing/domain"
	"database/sql"
	"strconv"
	"strings"
	"time"
)

type LoanRepository struct {
	loans []domain.Loan
	loan  domain.Loan
}

func NewLoanRepository() *LoanRepository {
	var domain []domain.Loan
	return &LoanRepository{loans: domain}
}

func (r *LoanRepository) Create(loan domain.Loan) (int64, error) {
	query := `INSERT INTO customer_loans (
		customer_id,
		interest_rate,
		installment,
		term_in_week,
		last_payment_term,
		principal,
		outstanding_loan,
		delinquent,
		created_at,
		updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	row, err := config.DB.Exec(query, loan.CustomerId, loan.InterestRate, loan.Installment, loan.TermInWeek, loan.LastPaymentTerm, loan.Principal, loan.OutstandingLoan, loan.Delinquent, loan.CreatedAt, loan.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *LoanRepository) CreateBillingSchedule(loan domain.Loan, customerLoanId int64) error {
	var billingSchedules []domain.BillingSchedule
	for week := 1; week <= int(loan.TermInWeek); week++ {
		billingSchedules = append(billingSchedules, domain.BillingSchedule{
			CustomerId:     loan.CustomerId,
			CustomerLoanId: customerLoanId,
			PaymentTerm:    int32(week),
			StartDate:      time.Now().UTC().Add(time.Duration(week*7-7) * 24 * time.Hour),
			DueDate:        time.Now().UTC().Add(time.Duration(week*7) * 24 * time.Hour),
		})
	}

	stmt, err := config.DB.Prepare("INSERT INTO customer_billing_schedules (customer_id, customer_loan_id, payment_term, start_date, due_date) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range billingSchedules {
		_, err := stmt.Exec(item.CustomerId, item.CustomerLoanId, item.PaymentTerm, item.StartDate, item.DueDate)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *LoanRepository) MakePayment(payment domain.Payment, lastPaymentTerm int, obligation int) (int, error) {
	values := []string{}
	for i := 1; i <= obligation; i++ {
		values = append(values, strconv.Itoa(lastPaymentTerm+i))
	}

	query := `UPDATE customer_billing_schedules  SET paid_at = ? WHERE customer_id = ? AND customer_loan_id = ? AND payment_term IN (` + strings.Join(values, ",") + `)`
	_, err := config.DB.Exec(query, payment.PaymentDate, payment.CustomerId, payment.CustomerLoanId)
	if err != nil {
		return 0, err
	}

	return lastPaymentTerm + obligation, nil
}

func (r *LoanRepository) CountObligation(paymentDate string, customerId int64, customerLoanId int64, startPaymentTerm int32) (int, error) {
	const layout = "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, paymentDate)
	if err != nil {
		return 0, err
	}

	var count int
	query := `SELECT  COUNT(*) AS count FROM customer_billing_schedules WHERE customer_id = ? AND customer_loan_id = ? AND start_date <= ? AND paid_at IS NULL;`
	row := config.DB.QueryRow(query, customerId, customerLoanId, parsedTime)
	if err := row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return count, nil
}

func (r *LoanRepository) GetCustomerLoan(id int64, customerId int64) (domain.Loan, error) {
	err := config.DB.QueryRow(`SELECT 
		id, 
		customer_id, 
		last_payment_term, 
		interest_rate, 
		term_in_week,  
		principal,
		installment,
		outstanding_loan,
		delinquent,
		created_at,
		updated_at
		FROM customer_loans WHERE id = ? AND customer_id = ?`, id, customerId).Scan(
		&r.loan.Id,
		&r.loan.CustomerId,
		&r.loan.LastPaymentTerm,
		&r.loan.InterestRate,
		&r.loan.TermInWeek,
		&r.loan.Principal,
		&r.loan.Installment,
		&r.loan.OutstandingLoan,
		&r.loan.Delinquent,
		&r.loan.CreatedAt,
		&r.loan.UpdatedAt,
	)

	if err != nil {
		return r.loan, err
	}

	return r.loan, nil
}

func (r *LoanRepository) AdjustCustomerLoan(loan domain.Loan) error {
	query := `UPDATE customer_loans  SET outstanding_loan = ?, last_payment_term = ?, delinquent = ? WHERE id = ?`
	_, err := config.DB.Exec(query, loan.OutstandingLoan, loan.LastPaymentTerm, loan.Delinquent, loan.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *LoanRepository) AdjustCustomerDelinquent(loan domain.Loan) error {
	query := `UPDATE customer_loans  SET delinquent = ? WHERE id = ?`
	_, err := config.DB.Exec(query, loan.Delinquent, loan.Id)
	if err != nil {
		return err
	}

	return nil
}
