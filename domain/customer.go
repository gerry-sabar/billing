package domain

import "time"

type Customer struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Delinquent  int       `json:"delinquent,omitempty" `
	Outstanding string    `json:"outstanding"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (Customer, error)
	Update(int, Customer) error
	Delete(int) error
	Create(customer Customer) (int64, error)
}
