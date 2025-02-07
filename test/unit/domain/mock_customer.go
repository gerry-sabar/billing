// Code generated by mockery v2.43.2. DO NOT EDIT.

package domain

import (
	domain "billing/domain"

	mock "github.com/stretchr/testify/mock"
)

// CustomerRepository is an autogenerated mock type for the CustomerRepository type
type CustomerRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: customer
func (_m *CustomerRepository) Create(customer domain.Customer) (int64, error) {
	ret := _m.Called(customer)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Customer) (int64, error)); ok {
		return rf(customer)
	}
	if rf, ok := ret.Get(0).(func(domain.Customer) int64); ok {
		r0 = rf(customer)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(domain.Customer) error); ok {
		r1 = rf(customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *CustomerRepository) Delete(_a0 int) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *CustomerRepository) GetAll() ([]domain.Customer, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []domain.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Customer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: _a0
func (_m *CustomerRepository) GetById(_a0 int) (domain.Customer, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 domain.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (domain.Customer, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) domain.Customer); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.Customer)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *CustomerRepository) Update(_a0 int, _a1 domain.Customer) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, domain.Customer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCustomerRepository creates a new instance of CustomerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCustomerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CustomerRepository {
	mock := &CustomerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
