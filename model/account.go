package main

// Transaction represents a single transaction
type Transaction struct {
	Amount             float64
	SourceAccount      string
	DestinationAccount string
}

// Account represents an account with transactions and a balance
type Account struct {
	AccountID    int64
	Transactions []Transaction
	Balance      float64
}

// NewAccount creates a new Account instance
func NewAccount(accountID int64, balance float64) *Account {
	return &Account{
		AccountID:    accountID,
		Transactions: []Transaction{},
		Balance:      balance,
	}
}

// AddTransaction adds a new transaction to the account
func (a *Account) AddTransaction(amount float64, sourceAccount, destinationAccount string) {
	transaction := Transaction{
		Amount:             amount,
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}
	a.Transactions = append(a.Transactions, transaction)
	a.Balance += amount
}
