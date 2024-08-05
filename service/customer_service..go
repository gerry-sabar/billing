package service

import "billing/domain"

type CustomerService interface {
	Create(data domain.Customer) (domain.Customer, error)
	GetById(id int) (domain.Customer, error)
	GetOutstanding(id int, loanId int) (domain.Customer, error)
	Update(id int, data domain.Customer) error
	GetAll() ([]domain.Customer, error)
	Delete(id int) error
}
