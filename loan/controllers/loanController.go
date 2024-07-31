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

// @Summary Create a new loan
// @Description Create a new loan with the provided details
// @Tags loans
// @Accept json
// @Produce json
// @Param loan body models.Loan true "Loan"
// @Success 201 {object} models.Loan
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /loans [post]
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

// @Summary Approve a loan
// @Description Approve a loan by its ID
// @Tags loans
// @Accept json
// @Produce json
// @Param id path int true "Loan ID"
// @Param approval body models.Approval true "Approval Details"
// @Success 200 {string} string "Loan approved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /loans/{id}/approve [put]
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

// @Summary Invest in a loan
// @Description Add investment to a loan
// @Tags loans
// @Accept json
// @Produce json
// @Param id path int true "Loan ID"
// @Param investment body models.Investment true "Investment Details"
// @Success 200 {string} string "Investment added successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /loans/{id}/invest [put]
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

// @Summary Disburse a loan
// @Description Disburse a loan by its ID
// @Tags loans
// @Accept json
// @Produce json
// @Param id path int true "Loan ID"
// @Param disbursement body models.Disbursement true "Disbursement Details"
// @Success 200 {string} string "Loan disbursed successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /loans/{id}/disburse [put]
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
