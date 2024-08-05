package server

import (
	"billing/handler"
	"billing/service"

	"github.com/gin-gonic/gin"
)

type RouterService struct {
	CustomerService service.CustomerService
	LoanService     service.LoanService
}

func SetupRouter(srv RouterService) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		customerHandler := handler.NewCustomerHandler(srv.CustomerService)
		v1.GET("/customers", customerHandler.GetAll())
		v1.POST("/customers", customerHandler.Create())
		v1.GET("/customers/:id", customerHandler.GetById())
		v1.PUT("/customers/:id", customerHandler.EditById())
		v1.DELETE("/customers/:id", customerHandler.Delete())
		v1.GET("/customers/:id/outstandings/:loanId", customerHandler.GetOutstanding())

		loanHandler := handler.NewLoanHandler(srv.LoanService)
		v1.POST("/loans", loanHandler.CreateLoan())
		v1.POST("/payments", loanHandler.MakePayment())
	}

	return router
}
