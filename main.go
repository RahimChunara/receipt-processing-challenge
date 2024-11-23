package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// in-memory stores
var receipts = make(map[string]Receipt)
var receiptPoints = make(map[string]int)

// handler for processing receipts
func processReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &receipt)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if !isValidReceipt(receipt) {
		http.Error(w, "Missing or invalid required fields", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	points := calculatePoints(receipt)
	receipts[id] = receipt
	receiptPoints[id] = points

	response := ProcessResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handler to retrieve points
func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := receiptPoints[id]
	if !exists {
		http.Error(w, "Receipt ID not found", http.StatusNotFound)
		return
	}

	response := PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", processReceiptHandler).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", getPointsHandler).Methods("GET")
	log.Println("Server started on :9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
