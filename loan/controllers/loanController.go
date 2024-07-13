package controllers

import (
	"loanEngine/common"
	"loanEngine/loan/models"
	loanService "loanEngine/loan/services"
	"loanEngine/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// The controller methods are updated to use the LoanService interface.
type LoanController struct {
	LoanService loanService.LoanService
}

func NewLoanController(service loanService.LoanService) *LoanController {
	return &LoanController{LoanService: service}
}

func (lc *LoanController) CreateLoan(c *gin.Context) {
	common.Logger.Info("CreateLoan-Started")
	defer common.Logger.Info("CreateLoan-End")
	var loan models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	common.Logger.Info("CreateLoan model", loan)
	if err := lc.LoanService.CreateLoan(&loan); err != nil {
		common.Logger.Error("err", err)
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, "Loan created successfully")
}

func (lc *LoanController) ApproveLoan(c *gin.Context) {
	common.Logger.Info("ApproveLoan-Started")
	defer common.Logger.Info("ApproveLoan-End")
	id, _ := strconv.Atoi(c.Param("id"))
	var approval models.Approval
	if err := c.ShouldBindJSON(&approval); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	common.Logger.Info("CreateLoan model", approval)
	if err := lc.LoanService.ApproveLoan(id, &approval); err != nil {
		common.Logger.Error("err", err)
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Loan approved successfully")
}

func (lc *LoanController) InvestLoan(c *gin.Context) {
	common.Logger.Info("InvestLoan-Started")
	defer common.Logger.Info("InvestLoan-End")
	id, _ := strconv.Atoi(c.Param("id"))
	var investment models.Investment
	if err := c.ShouldBindJSON(&investment); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	common.Logger.Info("CreateLoan model", investment)
	if err := lc.LoanService.InvestLoan(id, &investment); err != nil {
		common.Logger.Error("err", err)
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Investment added successfully")
}

func (lc *LoanController) DisburseLoan(c *gin.Context) {
	common.Logger.Info("DisburseLoan-Started")
	defer common.Logger.Info("DisburseLoan-End")
	id, _ := strconv.Atoi(c.Param("id"))
	var disbursement models.Disbursement
	if err := c.ShouldBindJSON(&disbursement); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	common.Logger.Info("CreateLoan model", disbursement)
	if err := lc.LoanService.DisburseLoan(id, &disbursement); err != nil {
		common.Logger.Error("err", err)
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Loan disbursed successfully")
}
