package models

type LoanState string

const (
	Proposed  LoanState = "proposed"
	Approved  LoanState = "approved"
	Invested  LoanState = "invested"
	Disbursed LoanState = "disbursed"
)

type Loan struct {
	ID              int        `json:"id" binding:"required"`
	BorrowerID      int        `json:"borrower_id" binding:"required"`
	PrincipalAmount float64    `json:"principal_amount" binding:"required"`
	Rate            float64    `json:"rate" binding:"required"`
	ROI             float64    `json:"roi" binding:"required"`
	AgreementLink   string     `json:"agreement_link"`
	State           LoanState  `json:"state" binding:"required"`
	Investments     []Investor `json:"investments"`
}

type Approval struct {
	PictureProof string `json:"picture_proof" binding:"required"`
	EmployeeID   int    `json:"employee_id" binding:"required"`
	ApprovalDate string `json:"approval_date" binding:"required"`
}

type Disbursement struct {
	AgreementLetter  string `json:"agreement_letter" binding:"required"`
	EmployeeID       int    `json:"employee_id" binding:"required"`
	DisbursementDate string `json:"disbursement_date" binding:"required"`
}

// Investor represents an investor in the loan system.
type Investor struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Email  string  `json:"email"`
}

type Investment struct {
	Investors []Investor `json:"investors"`
}

// ErrorResponse represents a standard error response format
type ErrorResponse struct {
	Error string `json:"error"`
}
