package service

import "billing/domain"

type LoanService interface {
	CreateLoan(data domain.LoanRequest) (domain.LoanRequest, error)
	MakePayment(data domain.PaymentRequest) error
	FormatCurrencyRequest(data domain.LoanRequest) (domain.Loan, error)
	FormatCurrencyResponse(data domain.Loan) domain.LoanRequest
}
