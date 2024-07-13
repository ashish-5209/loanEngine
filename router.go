package main

import (
	"loanEngine/loan/controllers"
	repository "loanEngine/loan/repositories"
	service "loanEngine/loan/services"

	"github.com/gin-gonic/gin"
)

// AddLoanRoutes adds loan routes(/loans/...).
func AddLoanRoutes(router *gin.Engine) {
	loanRepo := repository.NewInMemoryLoanRepository()
	loanService := service.NewLoanService(loanRepo)
	loanController := controllers.NewLoanController(loanService)
	api := router.Group("/api/v1")
	{
		api.POST("/loans", loanController.CreateLoan)
		api.PUT("/loans/:id/approve", loanController.ApproveLoan)
		api.PUT("/loans/:id/invest", loanController.InvestLoan)
		api.PUT("/loans/:id/disburse", loanController.DisburseLoan)
	}

}
