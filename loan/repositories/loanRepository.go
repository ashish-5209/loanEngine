package repository

import (
	"errors"
	"loanEngine/common"
	"loanEngine/loan/models"

	"go.uber.org/zap"
)

// The InMemoryLoanRepository implements the LoanRepository interface using an in-memory data structure
type InMemoryLoanRepository struct {
	loans       []models.Loan
	investments map[int][]models.Investment
}

func NewInMemoryLoanRepository() LoanRepository {
	return &InMemoryLoanRepository{
		loans:       []models.Loan{},
		investments: make(map[int][]models.Investment),
	}
}

func (r *InMemoryLoanRepository) CreateLoan(loan *models.Loan) error {
	loan.ID = len(r.loans) + 1
	r.loans = append(r.loans, *loan)
	common.Logger.Info("Loan created", zap.Int("loanID", loan.ID))
	common.Logger.Info("value", r)
	return nil
}

func (r *InMemoryLoanRepository) GetLoanByID(id int) (*models.Loan, error) {
	for _, loan := range r.loans {
		if loan.ID == id {
			common.Logger.Info("value", r)
			common.Logger.Info("Loan found", zap.Int("loanID", id))
			return &loan, nil
		}
	}
	common.Logger.Error("Loan not found", zap.Int("loanID", id))
	return nil, errors.New("loan not found")
}

func (r *InMemoryLoanRepository) UpdateLoan(loan *models.Loan) error {
	for i, l := range r.loans {
		if l.ID == loan.ID {
			r.loans[i] = *loan
			common.Logger.Info("Loan updated", zap.Int("loanID", loan.ID))
			common.Logger.Info("value", r)
			return nil
		}
	}
	common.Logger.Error("Loan not found for update", zap.Int("loanID", loan.ID))
	return errors.New("loan not found")
}

func (r *InMemoryLoanRepository) AddInvestment(loanID int, investment *models.Investment) error {
	r.investments[loanID] = append(r.investments[loanID], *investment)
	common.Logger.Info("Investment added", zap.Int("loanID", loanID))
	common.Logger.Info("value", r)
	return nil
}

func (r *InMemoryLoanRepository) GetTotalInvestedAmount(loanID int) float64 {
	total := 0.0
	for _, investment := range r.investments[loanID] {
		for _, val := range investment.Investors {
			total += val.Amount
		}
	}
	common.Logger.Info("Total invested amount", zap.Float64("total", total))
	return total
}

func (r *InMemoryLoanRepository) GetInvestorDetail() []models.Investor {
	var investors []models.Investor
	for _, loan := range r.loans {
		for _, investment := range r.investments[loan.ID] {
			for _, val := range investment.Investors {
				investors = append(investors, models.Investor{ID: val.ID, Email: val.Email})
			}
		}
	}
	common.Logger.Info("value", r)
	common.Logger.Info("Investors details retrieved", zap.Int("totalInvestors", len(investors)))
	return investors
}
