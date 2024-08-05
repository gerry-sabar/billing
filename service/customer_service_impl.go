package service

import (
	"billing/domain"
	"time"

	"github.com/shopspring/decimal"
)

type CustomerServiceImpl struct {
	CustomerRepository domain.CustomerRepository
	LoanRepository     domain.LoanRepository
}

func NewCustomerServiceImpl(customerRepository domain.CustomerRepository, loanRepository domain.LoanRepository) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		LoanRepository:     loanRepository,
	}
}

func (s *CustomerServiceImpl) Create(data domain.Customer) (domain.Customer, error) {
	customer := domain.Customer{
		Name:  data.Name,
		Email: data.Email,
	}

	id, err := s.CustomerRepository.Create(customer)
	if err != nil {
		return domain.Customer{}, err
	}

	customer.Id = id
	return customer, nil
}

func (s *CustomerServiceImpl) GetAll() ([]domain.Customer, error) {
	customers, err := s.CustomerRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s *CustomerServiceImpl) GetOutstanding(id int, loanId int) (domain.Customer, error) {
	customer, err := s.CustomerRepository.GetById(id)
	if err != nil {
		return domain.Customer{}, err
	}

	now := time.Now().UTC()
	formattedTime := now.Format("2006-01-02 15:04:05")
	customerLoan, err := s.LoanRepository.GetCustomerLoan(int64(loanId), customer.Id)
	if err != nil {
		return domain.Customer{}, err
	}

	customerObligation, err := s.LoanRepository.CountObligation(formattedTime, customer.Id, int64(loanId), customerLoan.LastPaymentTerm+1)
	if err != nil {
		return domain.Customer{}, err
	}

	if customerObligation > 2 {
		customerLoan.Delinquent = true
		customer.Delinquent = 1
		s.LoanRepository.AdjustCustomerDelinquent(customerLoan)
	}

	outstandingLoanDecimal := decimal.NewFromUint64(customerLoan.Installment).Mul(decimal.NewFromInt(int64(customerObligation))).Div(decimal.NewFromInt(100)).Round(2)
	customer.Outstanding = outstandingLoanDecimal.StringFixed(2)
	return customer, nil

}

func (s *CustomerServiceImpl) GetById(id int) (domain.Customer, error) {
	customer, err := s.CustomerRepository.GetById(id)
	if err != nil {
		return domain.Customer{}, err
	}

	return customer, nil
}

func (s *CustomerServiceImpl) Update(id int, data domain.Customer) error {
	_, err := s.CustomerRepository.GetById(id)
	if err != nil {
		return err
	}

	customer := domain.Customer{
		Name:  data.Name,
		Email: data.Email,
	}

	err = s.CustomerRepository.Update(id, customer)
	return err
}

func (s *CustomerServiceImpl) Delete(id int) error {
	_, err := s.CustomerRepository.GetById(id)
	if err != nil {
		return err
	}

	err = s.CustomerRepository.Delete(id)
	return err
}
