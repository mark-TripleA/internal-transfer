# Internal Transfer

## Installation

```bash
git clone https://github.com/mark-TripleA/internal-transfer.git
```

Navigate to the project directory: `cd internal-transfer`

Install the go modules: `go mod tidy`

Build the project: `go build`

Run the project: `go run main.go`

Sample request to create an account
```curl -X POST -H "Content-Type: application/json" -d '{"account_id": 123,"balance": 100.23344}' http://localhost:8080/accounts```
