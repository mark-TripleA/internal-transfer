package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Create a new synchronized map
var accounts sync.Map

// Transaction represents a single transaction
type Transaction struct {
	Amount             float64 `json:"amount"`
	SourceAccount      int64   `json:"source_account_id"`
	DestinationAccount int64   `json:"destination_account_id"`
}

// Account represents an account with transactions and a balance
type Account struct {
	AccountID    int64         `json:"account_id"`
	Transactions []Transaction `json:"transactions"`
	Balance      float64       `json:"balance"`
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
func (a *Account) AddTransaction(amount float64, sourceAccount, destinationAccount int64) *Transaction {
	transaction := Transaction{
		Amount:             amount,
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}
	a.Transactions = append(a.Transactions, transaction)
	if a.AccountID == sourceAccount {
		a.Balance -= amount
	} else if a.AccountID == destinationAccount {
		a.Balance += amount
	}
	return &transaction
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var newMessage Account
	json.NewDecoder(r.Body).Decode(&newMessage)

	account := NewAccount(newMessage.AccountID, newMessage.Balance)
	accounts.Store(newMessage.AccountID, account)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	var newMessage Account
	json.NewDecoder(r.Body).Decode(&newMessage)

	account, ok := accounts.Load(newMessage.AccountID)
	if !ok {
		// If account not found, return 404 Not Found
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Account with ID %v not found", newMessage.AccountID)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	json.NewDecoder(r.Body).Decode(&transaction)

	if account, ok := accounts.Load(transaction.DestinationAccount); ok {

		transaction := account.(*Account).AddTransaction(transaction.Amount, transaction.SourceAccount, transaction.DestinationAccount)

		// If account found, return the account details
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transaction)

	}

	if account, ok := accounts.Load(transaction.SourceAccount); ok {

		transaction := account.(*Account).AddTransaction(transaction.Amount, transaction.SourceAccount, transaction.DestinationAccount)

		// If account found, return the account details
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transaction)

	} else {
		// If account not found, return 404 Not Found
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Account with ID %s not found", transaction.DestinationAccount)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func handleRequests() {
	mux := http.NewServeMux()

	mux.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getAccount(w, r)
		} else if r.Method == http.MethodPost {
			createAccount(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/accounts/{account_id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the account ID from the request URL parameters
		// Parse the query parameters from the request
		queryParams := r.URL.Query()

		// Get the value of a specific query parameter
		accountID := queryParams.Get("account_id")

		if account, ok := accounts.Load(accountID); ok {
			// Account found, print its details
			fmt.Println("Account found:")
			// If account found, return the account details
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(account)

		} else {
			// If account not found, return 404 Not Found
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Account with ID %s not found", accountID)
			return
		}

	})
	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createTransaction(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {
	fmt.Println("API Started")
	handleRequests()
}
