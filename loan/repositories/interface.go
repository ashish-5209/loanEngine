package repository

import "loanEngine/loan/models"

// The LoanRepository interface defines the methods for the repository. This allows for different implementations (e.g., in-memory, database)
type LoanRepository interface {
	CreateLoan(loan *models.Loan) error
	GetLoanByID(id int) (*models.Loan, error)
	UpdateLoan(loan *models.Loan) error
	AddInvestment(loanID int, investment *models.Investment) error
	GetTotalInvestedAmount(loanID int) float64
	GetInvestorDetail() []models.Investor
}
