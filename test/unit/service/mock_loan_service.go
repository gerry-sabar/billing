// Code generated by mockery v2.43.2. DO NOT EDIT.

package service

import (
	domain "billing/domain"

	mock "github.com/stretchr/testify/mock"
)

// LoanService is an autogenerated mock type for the LoanService type
type LoanService struct {
	mock.Mock
}

// CreateLoan provides a mock function with given fields: data
func (_m *LoanService) CreateLoan(data domain.LoanRequest) (domain.LoanRequest, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for CreateLoan")
	}

	var r0 domain.LoanRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.LoanRequest) (domain.LoanRequest, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(domain.LoanRequest) domain.LoanRequest); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(domain.LoanRequest)
	}

	if rf, ok := ret.Get(1).(func(domain.LoanRequest) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormatCurrencyRequest provides a mock function with given fields: data
func (_m *LoanService) FormatCurrencyRequest(data domain.LoanRequest) (domain.Loan, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for FormatCurrencyRequest")
	}

	var r0 domain.Loan
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.LoanRequest) (domain.Loan, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(domain.LoanRequest) domain.Loan); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(domain.Loan)
	}

	if rf, ok := ret.Get(1).(func(domain.LoanRequest) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormatCurrencyResponse provides a mock function with given fields: data
func (_m *LoanService) FormatCurrencyResponse(data domain.Loan) domain.LoanRequest {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for FormatCurrencyResponse")
	}

	var r0 domain.LoanRequest
	if rf, ok := ret.Get(0).(func(domain.Loan) domain.LoanRequest); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(domain.LoanRequest)
	}

	return r0
}

// MakePayment provides a mock function with given fields: data
func (_m *LoanService) MakePayment(data domain.PaymentRequest) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for MakePayment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.PaymentRequest) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLoanService creates a new instance of LoanService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoanService(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoanService {
	mock := &LoanService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
