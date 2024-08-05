package domain

import (
	"time"
)

type Loan struct {
	Id              int64     `json:"id"`
	CustomerId      int64     `json:"customer_id" binding:"required"`
	InterestRate    float32   `json:"interest_rate" binding:"required"`
	TermInWeek      int32     `json:"term_in_week" binding:"required"`
	LastPaymentTerm int32     `json:"last_payment_term" binding:"required"`
	Principal       uint64    `json:"principal" binding:"required"`
	Installment     uint64    `json:"installment" binding:"required"`
	OutstandingLoan uint64    `json:"outstanding_loan" binding:"required"`
	Delinquent      bool      `json:"delinquent"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type LoanRequest struct {
	Id              int64     `json:"id"`
	CustomerId      int64     `json:"customer_id" binding:"required"`
	InterestRate    float32   `json:"interest_rate" binding:"required"`
	TermInWeek      int32     `json:"term_in_week" binding:"required"`
	LastPaymentTerm int32     `json:"last_payment_term"`
	Principal       string    `json:"principal" binding:"required"`
	Installment     string    `json:"installment"`
	OutstandingLoan string    `json:"outstanding_loan"`
	Delinquent      bool      `json:"delinquent"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type LoanRepository interface {
	Create(loan Loan) (int64, error)
	MakePayment(payment Payment, lastPaymentTerm int, obligation int) (int, error)
	CreateBillingSchedule(loan Loan, customerLoanId int64) error
	CountObligation(paymentDate string, customerId int64, customerLoanId int64, startPyamentTerm int32) (int, error)
	GetCustomerLoan(id int64, customerId int64) (Loan, error)
	AdjustCustomerLoan(loan Loan) error
	AdjustCustomerDelinquent(loan Loan) error
}
