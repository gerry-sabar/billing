// Code generated by mockery v2.43.2. DO NOT EDIT.

package domain

import (
	domain "billing/domain"

	mock "github.com/stretchr/testify/mock"
)

// LoanRepository is an autogenerated mock type for the LoanRepository type
type LoanRepository struct {
	mock.Mock
}

// AdjustCustomerLoan provides a mock function with given fields: loan
func (_m *LoanRepository) AdjustCustomerLoan(loan domain.Loan) error {
	ret := _m.Called(loan)

	if len(ret) == 0 {
		panic("no return value specified for AdjustCustomerLoan")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Loan) error); ok {
		r0 = rf(loan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountObligation provides a mock function with given fields: paymentDate, customerId, customerLoanId, startPyamentTerm
func (_m *LoanRepository) CountObligation(paymentDate string, customerId int64, customerLoanId int64, startPyamentTerm int32) (int, error) {
	ret := _m.Called(paymentDate, customerId, customerLoanId, startPyamentTerm)

	if len(ret) == 0 {
		panic("no return value specified for CountObligation")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64, int64, int32) (int, error)); ok {
		return rf(paymentDate, customerId, customerLoanId, startPyamentTerm)
	}
	if rf, ok := ret.Get(0).(func(string, int64, int64, int32) int); ok {
		r0 = rf(paymentDate, customerId, customerLoanId, startPyamentTerm)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, int64, int64, int32) error); ok {
		r1 = rf(paymentDate, customerId, customerLoanId, startPyamentTerm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: loan
func (_m *LoanRepository) Create(loan domain.Loan) (int64, error) {
	ret := _m.Called(loan)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Loan) (int64, error)); ok {
		return rf(loan)
	}
	if rf, ok := ret.Get(0).(func(domain.Loan) int64); ok {
		r0 = rf(loan)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(domain.Loan) error); ok {
		r1 = rf(loan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBillingSchedule provides a mock function with given fields: loan, customerLoanId
func (_m *LoanRepository) CreateBillingSchedule(loan domain.Loan, customerLoanId int64) error {
	ret := _m.Called(loan, customerLoanId)

	if len(ret) == 0 {
		panic("no return value specified for CreateBillingSchedule")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Loan, int64) error); ok {
		r0 = rf(loan, customerLoanId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCustomerLoan provides a mock function with given fields: id, customerId
func (_m *LoanRepository) GetCustomerLoan(id int64, customerId int64) (domain.Loan, error) {
	ret := _m.Called(id, customerId)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomerLoan")
	}

	var r0 domain.Loan
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64) (domain.Loan, error)); ok {
		return rf(id, customerId)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) domain.Loan); ok {
		r0 = rf(id, customerId)
	} else {
		r0 = ret.Get(0).(domain.Loan)
	}

	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(id, customerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakePayment provides a mock function with given fields: payment, lastPaymentTerm, obligation
func (_m *LoanRepository) MakePayment(payment domain.Payment, lastPaymentTerm int, obligation int) (int, error) {
	ret := _m.Called(payment, lastPaymentTerm, obligation)

	if len(ret) == 0 {
		panic("no return value specified for MakePayment")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Payment, int, int) (int, error)); ok {
		return rf(payment, lastPaymentTerm, obligation)
	}
	if rf, ok := ret.Get(0).(func(domain.Payment, int, int) int); ok {
		r0 = rf(payment, lastPaymentTerm, obligation)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(domain.Payment, int, int) error); ok {
		r1 = rf(payment, lastPaymentTerm, obligation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLoanRepository creates a new instance of LoanRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoanRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoanRepository {
	mock := &LoanRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
