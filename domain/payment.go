package domain

type Payment struct {
	Id             int64  `json:"id"`
	CustomerId     int64  `json:"customer_id" binding:"required"`
	CustomerLoanId int64  `json:"customer_loan_id" binding:"required"`
	PaymentDate    string `json:"payment_date" binding:"required"`
	Amount         uint64 `json:"amount" binding:"required"`
}

type PaymentRequest struct {
	Id             int64  `json:"id"`
	CustomerId     int64  `json:"customer_id" binding:"required"`
	CustomerLoanId int64  `json:"customer_loan_id" binding:"required"`
	PaymentDate    string `json:"payment_date"`
	Amount         string `json:"amount" binding:"required"`
}
