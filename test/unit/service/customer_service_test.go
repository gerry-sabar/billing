package service

import (
	domain "billing/domain"
	srv "billing/service"
	domainTest "billing/test/unit/domain"
	"errors"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllSuccess(t *testing.T) {
	expected := []domain.Customer{
		{
			Id:        1,
			Name:      "first test",
			Email:     "test@test.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        2,
			Name:      "second test",
			Email:     "test2@test.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("GetAll").Return(expected, nil)
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}
	result, err := service.GetAll()
	assert.Equal(t, expected, result)
	assert.NoError(t, err)
}

func TestGetAllFail(t *testing.T) {
	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("GetAll").Return(nil, errors.New("failed to get customers"))
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	todos, err := service.GetAll()
	assert.Error(t, err)
	assert.Equal(t, todos, []domain.Customer(nil))

}

func TestCreateSuccess(t *testing.T) {
	customer := domain.Customer{
		Name:  "test",
		Email: "test@test.com",
	}

	expected := domain.Customer{
		Id:    int64(1),
		Name:  "test",
		Email: "test@test.com",
	}

	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("Create", customer).Return(int64(1), nil)
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	result, err := service.Create(customer)
	assert.Equal(t, expected, result)
	assert.NoError(t, err)
}

func TestCreateFail(t *testing.T) {
	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("Create", domain.Customer{}).Return(int64(0), errors.New("failed to create customer"))
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	result, err := service.Create(domain.Customer{})
	assert.Equal(t, domain.Customer{}, result)
	assert.Error(t, err)
}
func TestUpdateSuccess(t *testing.T) {
	customer := domain.Customer{
		Id:        1,
		Name:      "test edit",
		Email:     "edit@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	update := domain.Customer{
		Name:  "test edit",
		Email: "edit@test.com",
	}

	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("GetById", 1).Return(customer, nil)
	mocking.On("Update", 1, update).Return(nil)
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	err := service.Update(1, update)
	assert.NoError(t, err)
}

func TestUpdateCustomerNotFound(t *testing.T) {
	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("GetById", 1).Return(domain.Customer{}, errors.New("customer is not found"))
	update := domain.Customer{
		Name:  "test edit",
		Email: "edit@test.com",
	}
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	err := service.Update(1, update)
	assert.Error(t, err)
	assert.Equal(t, "customer is not found", err.Error())
}

func TestUpdateFail(t *testing.T) {
	customer := domain.Customer{
		Id:        1,
		Name:      "test",
		Email:     "test@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	update := domain.Customer{
		Name:  "edit customer",
		Email: "edit@test.com",
	}

	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("GetById", 1).Return(customer, nil)
	mocking.On("Update", 1, update).Return(errors.New("failed to update customer"))
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	err := service.Update(1, update)
	assert.Error(t, err)
	assert.Equal(t, "failed to update customer", err.Error())
}

func TestDeleteSuccess(t *testing.T) {
	customer := domain.Customer{
		Id:        1,
		Name:      "test",
		Email:     "test@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("GetById", 1).Return(customer, nil)
	mocking.On("Delete", 1).Return(nil)
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	err := service.Delete(1)
	assert.NoError(t, err)
}

func TestDeleteCustomerNotFound(t *testing.T) {
	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("GetById", 1).Return(domain.Customer{}, errors.New("customer is not found"))
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	err := service.Delete(1)
	assert.Error(t, err)
	assert.Equal(t, "customer is not found", err.Error())
}

func TestDeleteFail(t *testing.T) {
	customer := domain.Customer{
		Id:        1,
		Name:      "test",
		Email:     "test@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mocking := domainTest.NewCustomerRepository(t)
	mocking.On("GetById", 1).Return(customer, nil)
	mocking.On("Delete", 1).Return(errors.New("failed to delete customer"))
	service := &srv.CustomerServiceImpl{
		CustomerRepository: mocking,
	}

	err := service.Delete(1)
	assert.Error(t, err)
	assert.Equal(t, "failed to delete customer", err.Error())
}
