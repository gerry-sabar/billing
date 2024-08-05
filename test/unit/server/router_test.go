package server

import (
	"billing/domain"
	"billing/server"
	"billing/test/unit/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	mocking := service.NewCustomerService(t)
	expected := []domain.Customer{
		{
			Id:          1,
			Name:        "first test",
			Email:       "test@test.com",
			Delinquent:  0,
			Outstanding: "",
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		},
		{
			Id:          2,
			Name:        "second test",
			Email:       "test2@test.com",
			Delinquent:  0,
			Outstanding: "",
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		},
	}
	RouterService := server.RouterService{
		CustomerService: mocking,
	}

	mocking.On("GetAll").Return(expected, nil)
	router := server.SetupRouter(RouterService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/customers", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string][]domain.Customer
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, response["data"][0], expected[0])
}

func TestCreate(t *testing.T) {
	mocking := service.NewCustomerService(t)
	now := time.Now().UTC()
	expected := domain.Customer{
		Name:  "test",
		Email: "test@test.com",
	}

	result := domain.Customer{
		Id:          1,
		Name:        "test",
		Email:       "test@test.com",
		Delinquent:  0,
		Outstanding: "",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	RouterService := server.RouterService{
		CustomerService: mocking,
	}

	body := `{"Name":"test", "Email":"test@test.com"}`
	mocking.On("Create", expected).Return(result, nil)

	router := server.SetupRouter(RouterService)
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/v1/customers", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]domain.Customer
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, response["data"], result)
}

func TestGetById(t *testing.T) {
	mocking := service.NewCustomerService(t)

	expected := domain.Customer{
		Id:          1,
		Name:        "first test",
		Email:       "test@test.com",
		Delinquent:  0,
		Outstanding: "",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	RouterService := server.RouterService{
		CustomerService: mocking,
	}

	mocking.On("GetById", 1).Return(expected, nil)
	router := server.SetupRouter(RouterService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/customers/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]domain.Customer
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, response["data"], expected)
}

func TestEditById(t *testing.T) {
	mocking := service.NewCustomerService(t)
	expected := domain.Customer{
		Name:  "test",
		Email: "test@test.com",
	}

	RouterService := server.RouterService{
		CustomerService: mocking,
	}

	body := `{"Name":"test", "Email":"test@test.com"}`
	mocking.On("Update", 1, expected).Return(nil, nil)

	router := server.SetupRouter(RouterService)
	w := httptest.NewRecorder()

	req := httptest.NewRequest("PUT", "/v1/customers/1", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, response["data"], "success")
}

func TestDelete(t *testing.T) {
	mocking := service.NewCustomerService(t)
	RouterService := server.RouterService{
		CustomerService: mocking,
	}

	mocking.On("Delete", 1).Return(nil, nil)

	router := server.SetupRouter(RouterService)
	w := httptest.NewRecorder()

	req := httptest.NewRequest("DELETE", "/v1/customers/1", nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetOutstanding(t *testing.T) {
	mocking := service.NewCustomerService(t)

	expected := domain.Customer{
		Id:          1,
		Name:        "first test",
		Email:       "test@test.com",
		Delinquent:  0,
		Outstanding: "110000.00",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	RouterService := server.RouterService{
		CustomerService: mocking,
	}

	mocking.On("GetOutstanding", 1, 1).Return(expected, nil)
	router := server.SetupRouter(RouterService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/customers/1/outstandings/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]domain.Customer
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, response["data"], expected)
}

func TestCreateLoan(t *testing.T) {
	mocking := service.NewLoanService(t)
	now := time.Now().UTC()

	request := domain.LoanRequest{
		CustomerId:   1,
		InterestRate: 0.1,
		TermInWeek:   50,
		Principal:    "5000000.00",
	}

	expected := domain.LoanRequest{
		Id:              1,
		CustomerId:      1,
		InterestRate:    0.1,
		TermInWeek:      50,
		LastPaymentTerm: 0,
		Principal:       "5000000.00",
		Installment:     "110000.00",
		OutstandingLoan: "5500000.00",
		Delinquent:      false,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	RouterService := server.RouterService{
		LoanService: mocking,
	}

	body := `{"customer_id":1, "interest_rate": 0.1, "term_in_week": 50, "principal": "5000000.00"}`
	mocking.On("CreateLoan", request).Return(expected, nil)

	router := server.SetupRouter(RouterService)
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/v1/loans", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]domain.LoanRequest
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, response["data"], expected)
}

func TestMakePayment(t *testing.T) {
	mocking := service.NewLoanService(t)

	request := domain.PaymentRequest{
		CustomerId:     1,
		CustomerLoanId: 1,
		Amount:         "220000.00",
	}

	RouterService := server.RouterService{
		LoanService: mocking,
	}

	body := `{"customer_id":1, "customer_loan_id": 1, "amount": "220000.00"}`
	mocking.On("MakePayment", request).Return(nil, nil)

	router := server.SetupRouter(RouterService)
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/v1/payments", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, response["data"], "success")

}
