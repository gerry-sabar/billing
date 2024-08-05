package repository

import (
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"billing/config"
	"billing/domain"
	"billing/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO customers (name, email, delinquent, created_at, updated_at) VALUES (?, ?, ?, now(), now())")).
		WithArgs("test", "test@test.com", 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	var repo repository.CustomerRepository
	var customer domain.Customer
	customer.Name = "test"
	customer.Email = "test@test.com"
	id, err := repo.Create(customer)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE customers SET name = ?, email = ? where id = ?")).
		WithArgs("update edit", "edit@edit.com", int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var repo repository.CustomerRepository
	var customer domain.Customer
	customer.Name = "update edit"
	customer.Email = "edit@edit.com"
	err = repo.Update(1, customer)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM customers WHERE id = ?")).
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var repo repository.CustomerRepository
	err = repo.Delete(1)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	config.DB = db
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer config.DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "createdAt", "updatedAt"}).AddRow(1, "test", "test@test.com", 0, time.Now(), time.Now())
	mock.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `name`, `email`, `created_at`, `updated_at` FROM customers WHERE id = ?")).
		WithArgs(1).
		WillReturnRows(rows)

	var repo repository.CustomerRepository
	customer, err := repo.GetById(1)

	assert.NoError(t, err)
	assert.Equal(t, "test", customer.Name)
	assert.Equal(t, "test@test.com", customer.Email)

}
