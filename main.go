package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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

type Message struct {
	Text string `json:"text"`
}

var messages []Message

func createAccount(w http.ResponseWriter, r *http.Request) {
	var newMessage Message
	json.NewDecoder(r.Body).Decode(&newMessage)
	messages = append(messages, newMessage)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMessage)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	var newMessage Message
	json.NewDecoder(r.Body).Decode(&newMessage)
	messages = append(messages, newMessage)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMessage)
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
	messages = []Message{
		Message{Text: "Welcome to the API!"},
	}
	handleRequests()
}
