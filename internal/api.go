package api

import (
	"encoding/json"
	"log"
	"net/http"
)

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
