# loanEngine
LoanEngine is a service designed to handle loan management with multiple states: proposed, approved, invested, and disbursed. This service provides a RESTful API to manage loans, handle state transitions, and notify investors via email.

# LOANENGINE

State Management: Loans can be in multiple states - proposed, approved, invested, and disbursed. Each state has specific rules and transitions.
Email Notifications: Once a loan is invested, all investors receive an email with a link to the agreement letter.
PDF Generation: Agreement letters are generated and stored as PDF files.
RESTful API: Exposes endpoints to create, update, and manage loans.
---
## Requirements
* Go1.22

---
## Dependecies
* loanengine_path: Point to the `loanengine` directory.
For exporting these variables, add the following command in ~/.zshrc
```
# loanengine config
export loanengine_path='directory/path/.../loanengine'
```

* Environment.env: Every env has a dedicated $Environment.env file (dev.env, pp.env, prod.env). \
For exporting these variables, add the following command in ~/.zshrc
```
# loanengine config
export $(xargs < ${loanengine_path}/credentials/dev.env)
```

---

## Build and Run
 build: ```go build -v -o ./bin/loanengine loanengine``` or ```make```(creates a binary named loanengine) \
 run: ```make run``` (executes the binary in dev environment, to run in different environment update the .env file path in Makefile) \
 release: Builds for Prod env and runs on Prod. Dependencies: ssh config named Amarthaloan

---
<br/>

## API Endpoints
1. Create Loan
    Endpoint: POST /api/v1/loans/
    Description: Create a new loan with the initial state proposed.
    Request Body:
        {
            "id": 1,
            "borrower_id": 1,
            "principal_amount": 10000,
            "rate": 5.0,
            "roi": 10.0,
            "agreement_link": "http://example.com/agreement.pdf",
            "state": "proposed"
        }

2. Approve Loan
    Endpoint: PUT /api/v1/loans/:id/approve
    Description: Approve a proposed loan.
    Request Body:
        {
            "picture_proof": "http://example.com/proof.jpg",
            "employee_id": 123,
            "approval_date": "2024-07-13"
        }

3. Invest in Loan
    Endpoint: PUT /api/v1/loans/:id/invest
    Description: Invest in an approved loan.
    Request Body:
        {
            "investors": [
                    {
                    "id": 1,
                    "amount": 5000,
                    "email": "investor1@example.com"
                },
                {
                    "id": 2,
                    "amount": 5000,
                    "email": "investor2@example.com"
                }
            ]
        }

4. Disburse Loan
    Endpoint: PUT /api/v1/loans/:id/disburse
    Description: Disburse an invested loan.
    Request Body:
        {
            "agreement_letter": "http://example.com/agreement.pdf",
            "employee_id": 456,
            "disbursement_date": "2024-07-13"
        }

## Docker
docker build -t loanengine:latest .
docker run -p 5002:5002 loanengine:latest

## Logging
Logs are stored in /var/log/loanEngine. Ensure this directory exists and has the correct permissions.

## Environment Variables
    agreements: Path where agreement PDFs are stored.
    MAILGUN_DOMAIN: Domain for Mailgun.
    MAILGUN_API_KEY: API key for Mailgun.
    MAILGUN_SENDER_EMAIL: Sender email for Mailgun.

---