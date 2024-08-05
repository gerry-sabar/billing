package repository

import (
	"billing/config"
	"billing/domain"
)

type CustomerRepository struct {
	customers []domain.Customer
	customer  domain.Customer
}

func NewCustomerRepository() *CustomerRepository {
	var domain []domain.Customer
	return &CustomerRepository{customers: domain}
}

func (r *CustomerRepository) Create(customer domain.Customer) (int64, error) {
	query := `INSERT INTO customers (name, email, created_at, updated_at) VALUES (?, ?, now(), now())`
	row, err := config.DB.Exec(query, customer.Name, customer.Email)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *CustomerRepository) Update(id int, customer domain.Customer) error {
	_, err := config.DB.Exec("UPDATE customers SET name = ?, email = ?, where id = ?", customer.Name, customer.Email, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM customers WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) GetById(id int) (domain.Customer, error) {
	err := config.DB.QueryRow("SELECT `id`, `name`, `email`, `created_at`, `updated_at` FROM customers WHERE id = ?", id).Scan(&r.customer.Id, &r.customer.Name, &r.customer.Email, &r.customer.CreatedAt, &r.customer.UpdatedAt)
	if err != nil {
		return r.customer, err
	}

	return r.customer, nil
}

func (r *CustomerRepository) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	rows, err := config.DB.Query("SELECT `id`, `name`, `email`, `created_at`, `updated_at` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer domain.Customer
		if err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}
