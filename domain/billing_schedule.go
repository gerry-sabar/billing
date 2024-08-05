package domain

import "time"

type BillingSchedule struct {
	Id             int64     `json:"id"`
	CustomerId     int64     `json:"customer_id" binding:"required"`
	CustomerLoanId int64     `json:"customer_load_id" binding:"required"`
	PaymentTerm    int32     `json:"payment_term" binding:"required"`
	StartDate      time.Time `json:"start_date" binding:"required"`
	DueDate        time.Time `json:"due_date" binding:"required"`
	PaidAt         time.Time `json:"paid_at"`
}
