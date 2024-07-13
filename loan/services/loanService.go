package services

import (
	"context"
	"errors"
	"fmt"
	"loanEngine/common"
	"loanEngine/loan/models"
	loanRepository "loanEngine/loan/repositories"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/mailgun/mailgun-go/v4"
	"go.uber.org/zap"
)

// the loanService struct implements these methods.
type loanService struct {
	repository loanRepository.LoanRepository
}

func NewLoanService(repo loanRepository.LoanRepository) LoanService {
	return &loanService{repository: repo}
}

func (s *loanService) CreateLoan(loan *models.Loan) error {
	loan.State = models.Proposed
	return s.repository.CreateLoan(loan)
}

func (s *loanService) ApproveLoan(id int, approval *models.Approval) error {
	loan, err := s.repository.GetLoanByID(id)
	if err != nil {
		common.Logger.Error("err", err)
		return err
	}
	if loan.State != models.Proposed {
		common.Logger.Error("err", errors.New("loan can only be approved from proposed state"))
		return errors.New("loan can only be approved from proposed state")
	}
	loan.State = models.Approved
	// Save approval information, update loan state in the repository
	return s.repository.UpdateLoan(loan)
}

func (s *loanService) InvestLoan(id int, investment *models.Investment) error {
	loan, err := s.repository.GetLoanByID(id)
	if err != nil {
		common.Logger.Error("err", err)
		return err
	}
	if loan.State != models.Approved {
		common.Logger.Error("err", errors.New("loan can only be invested in approved state"))
		return errors.New("loan can only be invested in approved state")
	}
	// Add investment, check if total invested amount equals principal amount
	amt := 0.0
	for _, val := range investment.Investors {
		if val.ID == id {
			amt += val.Amount
		}
	}
	if s.repository.GetTotalInvestedAmount(id)+amt > loan.PrincipalAmount {
		common.Logger.Error("err", errors.New("total invested amount cannot be bigger than principal amount"))
		return errors.New("total invested amount cannot be bigger than principal amount")
	}
	s.repository.AddInvestment(id, investment)
	if s.repository.GetTotalInvestedAmount(id) == loan.PrincipalAmount {
		loan.State = models.Invested
		investor := s.repository.GetInvestorDetail()

		// generate and send agreement letter to all investors
		go FinalizeInvestment(loan, investor)
		s.repository.UpdateLoan(loan)
	}
	return nil
}

// SendEmail sends an email using Mailgun
func SendEmail(to []string, subject, body string) {
	mg := mailgun.NewMailgun(mailgunDomain, mailgunAPIKey)

	message := mg.NewMessage(
		mailgunSender,
		subject,
		"",
		to...,
	)
	message.SetHtml(body)

	// Send the email
	ctx := context.Background()
	_, _, err := mg.Send(ctx, message)
	if err != nil {
		common.Logger.Error("Error sending email", zap.Error(err))

	}

	common.Logger.Info("Email sent successfully")
}

func (s *loanService) DisburseLoan(id int, disbursement *models.Disbursement) error {
	loan, err := s.repository.GetLoanByID(id)
	if err != nil {
		common.Logger.Error("err", err)
		return err
	}
	if loan.State != models.Invested {
		common.Logger.Error("err", errors.New("loan can only be disbursed from invested state"))
		return errors.New("loan can only be disbursed from invested state")
	}
	loan.State = models.Disbursed
	// Save disbursement information, update loan state in the repository
	return s.repository.UpdateLoan(loan)
}

// SendInvestmentNotification sends emails to all investors with the link to the agreement letter
func SendInvestmentNotification(investors []models.Investor, agreementLetterPath string) error {
	for _, investor := range investors {
		// Prepare email content
		subject := "Your Investment Agreement Letter"
		body := fmt.Sprintf(
			"Dear Investor,<br><br>Thank you for your investment. Please find your agreement letter at the following link:<br><a href=\"%s\">Download Agreement</a><br><br>Best regards,<br>Loan Engine",
			agreementLetterPath,
		)

		// Send email, used goroutines
		SendEmail([]string{investor.Email}, subject, body)
	}
	return nil
}

func FinalizeInvestment(loan *models.Loan, investors []models.Investor) error {

	// Generate agreement letter
	agreementLetterPath, err := generateAgreementLetter(loan)
	if err != nil {
		common.Logger.Error("err", err)
		return err
	}

	// Notify investors
	err = SendInvestmentNotification(investors, agreementLetterPath)
	if err != nil {
		common.Logger.Error("err", err)
		return err
	}

	return nil
}

func generateAgreementLetter(loan *models.Loan) (string, error) {
	directory := "./agreements"
	fileName := fmt.Sprintf("agreement_%d_%s.pdf", loan.ID, time.Now().Format("20060102150405"))
	filePath := fmt.Sprintf("%s/%s", directory, fileName)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			common.Logger.Error("Error creating agreements directory", zap.Error(err))
			return "", err
		}
	}

	// Initialize PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Write content to PDF
	pdf.Cell(0, 10, fmt.Sprintf("Agreement for Loan ID %d", loan.ID))
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Borrower ID: %d", loan.BorrowerID))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Principal Amount: %.2f", loan.PrincipalAmount))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Rate: %.2f%%", loan.Rate))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("ROI: %.2f%%", loan.ROI))
	pdf.Ln(10)
	pdf.MultiCell(0, 10, "This agreement confirms the terms of the loan provided. Please review the details carefully.", "", "", false)
	pdf.Ln(10)

	// Add the "Please sign" text
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Please sign below:")
	pdf.Ln(20)

	// Draw a line for the signature
	pdf.SetDrawColor(0, 0, 0) // Black color for the line
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
	pdf.Ln(10)
	pdf.Cell(0, 10, "Signature")

	// Save the PDF to file
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		common.Logger.Error("Error writing to agreement file", zap.Error(err))
		return "", err
	}

	// Log the success and return the file path
	common.Logger.Info("Agreement letter generated", zap.String("filePath", filePath))
	return filePath, nil
}
