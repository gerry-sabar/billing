package main

import (
	"billing/config"
	"billing/database"
	"billing/repository"
	"billing/server"
	"billing/service"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	config.LoadConfig()
	config.DB, err = database.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer config.DB.Close()

	LoanRepository := repository.NewLoanRepository()
	loanService := service.NewLoanServiceImpl(LoanRepository)
	CustomerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerServiceImpl(CustomerRepository, LoanRepository)

	srv := server.RouterService{
		CustomerService: customerService,
		LoanService:     loanService,
	}

	r := server.SetupRouter(srv)
	r.Run(fmt.Sprintf(":%s", config.Cfg.APPPort))
}
