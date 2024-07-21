package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        int
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

var database = map[int]Customer{
	0: {
		ID:        1,
		Name:      "Mustafa",
		Role:      "Software Engineering",
		Email:     "m.almeer.swe@gmail.com",
		Phone:     "0543584464",
		Contacted: false,
	},
	1: {
		ID:        2,
		Name:      "Ahmed",
		Role:      "Software Engineering",
		Email:     "ahmed@gmail.com",
		Phone:     "unk4",
		Contacted: false,
	},
	2: {
		ID:        3,
		Name:      "GG",
		Role:      "Software Engineering",
		Email:     "m.GG.swe@gmail.com",
		Phone:     "da",
		Contacted: false,
	},
}

// Retrive all customers from DB
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database)
}

// Retrieve single customer from DB by {id}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL and convert to integer
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	// Lookup customer in database
	customer, exists := database[id]
	if !exists {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	// Encode and respond with customer data
	if err := json.NewEncoder(w).Encode(customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer Customer
	json.NewDecoder(r.Body).Decode(&newCustomer)

	key := len(database)
	newID := len(database) + 1
	newCustomer.ID = newID

	database[key] = newCustomer

	json.NewEncoder(w).Encode(newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	existingCustomer, exist := database[id]
	if !exist {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	var updatedCustomer Customer
	if err := json.NewDecoder(r.Body).Decode(&updatedCustomer); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Preserve the original ID
	updatedCustomer.ID = existingCustomer.ID

	database[id] = updatedCustomer

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedCustomer); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// define content type
	w.Header().Set("Content-Type", "application/json")

	// user request
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Error occure in deleteCustomer()", http.StatusBadRequest)
		return
	}

	for key, customer := range database {
		if customer.ID == id {
			delete(database, key)
			break
		}
	}

	json.NewEncoder(w).Encode(database)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
