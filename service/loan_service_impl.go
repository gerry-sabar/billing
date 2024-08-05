package service

import (
	"billing/config"
	"billing/domain"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

type LoanServiceImpl struct {
	CustomerRepository domain.CustomerRepository
	LoanRepository     domain.LoanRepository
}

func NewLoanServiceImpl(loanRepository domain.LoanRepository) *LoanServiceImpl {
	return &LoanServiceImpl{
		LoanRepository: loanRepository,
	}
}

func (s *LoanServiceImpl) FormatCurrencyRequest(data domain.LoanRequest) (domain.Loan, error) {
	var loan domain.Loan
	currencyDivider := decimal.NewFromInt(100)

	principal, err := decimal.NewFromString(data.Principal)
	if err != nil {
		return loan, err
	}

	principalDecimal := principal.Round(2).Mul(currencyDivider)
	interestRateDecimal := decimal.NewFromFloat(float64(data.InterestRate)).Round(3)
	interestDecimal := principalDecimal.Mul(interestRateDecimal).Mul(decimal.NewFromInt32(data.TermInWeek)).Div(decimal.NewFromInt(50))
	outstandingLoanDecimal := principalDecimal.Add(interestDecimal).Round(2)
	installment := outstandingLoanDecimal.Div(decimal.NewFromInt32(data.TermInWeek)).Round(2)

	loan = domain.Loan{
		CustomerId:      data.CustomerId,
		InterestRate:    float32(interestRateDecimal.InexactFloat64()),
		Installment:     installment.BigInt().Uint64(),
		TermInWeek:      data.TermInWeek,
		LastPaymentTerm: 0,
		Principal:       principalDecimal.BigInt().Uint64(),
		OutstandingLoan: outstandingLoanDecimal.BigInt().Uint64(),
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		Delinquent:      false,
	}

	return loan, nil
}

func (s *LoanServiceImpl) FormatCurrencyResponse(data domain.Loan) domain.LoanRequest {
	currencyDivider := decimal.NewFromInt(100)

	principalDecimal := decimal.NewFromUint64(data.Principal).Div(currencyDivider).Round(2)
	installment := decimal.NewFromUint64(data.Installment).Div(currencyDivider).Round(2)
	outstandingLoanDecimal := decimal.NewFromUint64(data.Installment).Div(currencyDivider).Round(2)

	loan := domain.LoanRequest{
		CustomerId:      data.CustomerId,
		InterestRate:    data.InterestRate,
		Installment:     installment.StringFixed(2),
		TermInWeek:      data.TermInWeek,
		LastPaymentTerm: 0,
		Principal:       principalDecimal.StringFixed(2),
		OutstandingLoan: outstandingLoanDecimal.StringFixed(2),
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		Delinquent:      false,
	}

	return loan
}

func (s *LoanServiceImpl) CreateLoan(data domain.LoanRequest) (domain.LoanRequest, error) {
	//one year is based on 50 weeks instead of 52 weeks
	loan, err := s.FormatCurrencyRequest(data)
	if err != nil {
		return domain.LoanRequest{}, err
	}

	tx, err := config.DB.Begin()

	customerLoanId, err := s.LoanRepository.Create(loan)
	if err != nil {
		return domain.LoanRequest{}, err
	}

	err = s.LoanRepository.CreateBillingSchedule(loan, customerLoanId)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	response := s.FormatCurrencyResponse(loan)
	response.Id = customerLoanId
	return response, nil
}

func (s *LoanServiceImpl) MakePayment(data domain.PaymentRequest) error {
	currencyDivider := decimal.NewFromInt(100)
	now := time.Now().UTC()
	formattedTime := now.Format("2006-01-02 15:04:05")
	amount, err := decimal.NewFromString(data.Amount)
	if err != nil {
		return err
	}

	amountDecimal := amount.Mul(currencyDivider)
	payment := domain.Payment{
		CustomerId:     data.CustomerId,
		CustomerLoanId: data.CustomerLoanId,
		Amount:         amountDecimal.BigInt().Uint64(),
		PaymentDate:    formattedTime,
	}

	tx, err := config.DB.Begin()

	loanDetail, err := s.LoanRepository.GetCustomerLoan(data.CustomerLoanId, data.CustomerId)
	if err != nil {
		return err
	}

	totalObligation, err := s.LoanRepository.CountObligation(formattedTime, data.CustomerId, data.CustomerLoanId, loanDetail.LastPaymentTerm+1)
	if err != nil {
		return err
	}

	totalInstallment := loanDetail.Installment * uint64(totalObligation)
	fmt.Println(fmt.Sprintf("Wrong amount, the installment you have to pay is: %d", totalInstallment/100))
	if payment.Amount != totalInstallment {
		return errors.New(fmt.Sprintf("Wrong amount, the installment you have to pay is: %.2d", totalInstallment/100))
	}

	lastPaymentTermPos, err := s.LoanRepository.MakePayment(payment, int(loanDetail.LastPaymentTerm), totalObligation)
	if err != nil {
		return err
	}

	outstanding := decimal.NewFromUint64(loanDetail.OutstandingLoan).Sub(amountDecimal)
	loanDetail.OutstandingLoan = outstanding.BigInt().Uint64()
	loanDetail.LastPaymentTerm = int32(lastPaymentTermPos)
	if totalObligation >= 3 {
		loanDetail.Delinquent = true
	}

	err = s.LoanRepository.AdjustCustomerLoan(loanDetail)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
