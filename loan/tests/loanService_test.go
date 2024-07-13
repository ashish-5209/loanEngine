package tests

import (
	"errors"
	"loanEngine/loan/models"
	repository "loanEngine/loan/repositories"
	"loanEngine/loan/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases are added using the testify package to verify the functionality of the service layer. Mock repositories are used to isolate the service layer for testing.
type MockLoanRepository struct {
	loans       []models.Loan
	investments map[int][]models.Investment
}

func NewMockLoanRepository() repository.LoanRepository {
	return &MockLoanRepository{
		loans:       []models.Loan{},
		investments: make(map[int][]models.Investment),
	}
}

func (r *MockLoanRepository) CreateLoan(loan *models.Loan) error {
	loan.ID = len(r.loans) + 1
	r.loans = append(r.loans, *loan)
	return nil
}

func (r *MockLoanRepository) GetLoanByID(id int) (*models.Loan, error) {
	for _, loan := range r.loans {
		if loan.ID == id {
			return &loan, nil
		}
	}
	return nil, errors.New("loan not found")
}

func (r *MockLoanRepository) UpdateLoan(loan *models.Loan) error {
	for i, l := range r.loans {
		if l.ID == loan.ID {
			r.loans[i] = *loan
			return nil
		}
	}
	return errors.New("loan not found")
}

func (r *MockLoanRepository) AddInvestment(loanID int, investment *models.Investment) error {
	r.investments[loanID] = append(r.investments[loanID], *investment)
	return nil
}

func (r *MockLoanRepository) GetTotalInvestedAmount(loanID int) float64 {
	var total float64
	for idx, investment := range r.investments[loanID] {
		total += investment.Investors[idx].Amount
	}
	return total
}

func (r *MockLoanRepository) GetInvestorDetail() []models.Investor {
	var investors []models.Investor
	for _, loan := range r.loans {
		for _, investment := range r.investments[loan.ID] {
			for _, val := range investment.Investors {
				investors = append(investors, models.Investor{ID: val.ID, Email: val.Email})
			}
		}
	}
	return investors
}

func TestCreateLoan(t *testing.T) {
	repo := NewMockLoanRepository()
	service := services.NewLoanService(repo)

	loan := models.Loan{
		BorrowerID:      1,
		PrincipalAmount: 10000,
		Rate:            5,
		ROI:             7,
		AgreementLink:   "",
		State:           models.Proposed,
	}

	err := service.CreateLoan(&loan)
	assert.Nil(t, err)
	assert.Equal(t, models.Proposed, loan.State)
	assert.Equal(t, 1, loan.ID)
}

func TestApproveLoan(t *testing.T) {
	repo := NewMockLoanRepository()
	service := services.NewLoanService(repo)

	loan := models.Loan{
		BorrowerID:      1,
		PrincipalAmount: 10000,
		Rate:            5,
		ROI:             7,
		AgreementLink:   "",
		State:           models.Proposed,
	}

	repo.CreateLoan(&loan)
	approval := models.Approval{
		PictureProof: "/Users/sheelaashish/Downloads/Ashish.jpeg",
		EmployeeID:   1,
		ApprovalDate: "2023-01-01",
	}

	err := service.ApproveLoan(loan.ID, &approval)
	assert.Nil(t, err)

	updatedLoan, _ := repo.GetLoanByID(loan.ID)
	assert.Equal(t, models.Approved, updatedLoan.State)
}

func TestInvestLoan(t *testing.T) {
	repo := NewMockLoanRepository()
	service := services.NewLoanService(repo)

	loan := models.Loan{
		BorrowerID:      1,
		PrincipalAmount: 10000,
		Rate:            5,
		ROI:             7,
		AgreementLink:   "",
		State:           models.Approved,
	}

	repo.CreateLoan(&loan)
	investors := []models.Investor{
		{ID: 1, Amount: 10000, Email: "ashishkumarsaw3@gmail.com"},
	}
	investment := models.Investment{
		Investors: investors,
	}

	err := service.InvestLoan(loan.ID, &investment)
	assert.Nil(t, err)

	updatedLoan, _ := repo.GetLoanByID(loan.ID)
	assert.Equal(t, models.Invested, updatedLoan.State)
}

func TestDisburseLoan(t *testing.T) {
	repo := NewMockLoanRepository()
	service := services.NewLoanService(repo)

	loan := models.Loan{
		BorrowerID:      1,
		PrincipalAmount: 10000,
		Rate:            5,
		ROI:             7,
		AgreementLink:   "",
		State:           models.Invested,
	}

	repo.CreateLoan(&loan)
	disbursement := models.Disbursement{
		AgreementLetter:  "agreement.pdf",
		EmployeeID:       1,
		DisbursementDate: "2023-01-01",
	}

	err := service.DisburseLoan(loan.ID, &disbursement)
	assert.Nil(t, err)

	updatedLoan, _ := repo.GetLoanByID(loan.ID)
	assert.Equal(t, models.Disbursed, updatedLoan.State)
}
