package services

import (
	"loanEngine/loan/models"
)

// The LoanService interface defines the service methods
type LoanService interface {
	CreateLoan(loan *models.Loan) error
	ApproveLoan(id int, approval *models.Approval) error
	InvestLoan(id int, investment *models.Investment) error
	DisburseLoan(id int, disbursement *models.Disbursement) error
}
