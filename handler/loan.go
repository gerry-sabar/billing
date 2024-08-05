package handler

import (
	"billing/domain"
	"billing/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ILoanHandler interface {
	Create() gin.HandlerFunc
	CreatePayment() gin.HandlerFunc
}

type LoanService struct {
	Loan domain.Loan
}

type LoanHandler struct {
	LoanService service.LoanService
}

func NewLoanHandler(loanService service.LoanService) *LoanHandler {
	return &LoanHandler{LoanService: loanService}
}

func (h *LoanHandler) CreateLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		loanReq := domain.LoanRequest{}
		if err := c.ShouldBind(&loanReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.LoanService.CreateLoan(loanReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func (h *LoanHandler) MakePayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		paymentReq := domain.PaymentRequest{}
		if err := c.ShouldBind(&paymentReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.LoanService.MakePayment(paymentReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
