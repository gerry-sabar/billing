package handler

import (
	"net/http"
	"strconv"

	"billing/domain"
	"billing/service"

	"github.com/gin-gonic/gin"
)

type ICustomerHandler interface {
	Delete() gin.HandlerFunc
	EditById() gin.HandlerFunc
	GetById() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	Create() gin.HandlerFunc
	GetOutstanding() gin.HandlerFunc
	// CreateLoan() gin.HandlerFunc
}

type CustomerService struct {
	Customer domain.Customer
}

type CustomerHandler struct {
	CustomerService service.CustomerService
	LoanService     service.LoanService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{CustomerService: customerService}
}

func (h *CustomerHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.CustomerService.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (h *CustomerHandler) EditById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customerReq := domain.Customer{}
		if err := c.ShouldBind(&customerReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.CustomerService.Update(id, customerReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}

func (h *CustomerHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customer, _ := h.CustomerService.GetById(id)
		c.JSON(http.StatusOK, gin.H{"data": customer})
	}
}

func (h *CustomerHandler) GetOutstanding() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		loanId, err := strconv.Atoi(c.Param("loanId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customer, err := h.CustomerService.GetOutstanding(id, loanId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": customer})
	}
}

func (h *CustomerHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		customers, _ := h.CustomerService.GetAll()
		c.JSON(http.StatusOK, gin.H{"data": customers})
	}
}

func (h *CustomerHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var customerReq domain.Customer
		if err := c.ShouldBind(&customerReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.CustomerService.Create(customerReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
