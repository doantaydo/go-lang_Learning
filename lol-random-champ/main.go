package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RequestData represents the data structure for the incoming request
type RequestData struct {
	UserInput string `json:"userInput"`
}

// ResponseData represents the data structure for the server response
type ResponseData struct {
	Message string `json:"message"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Only handle POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming JSON payload
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Your server logic using requestData.UserInput
	responseData := ResponseData{Message: "Hello from Golang server!"}

	// Convert the response data to JSON
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(responseJSON)
}

func main() {
	// Handle requests to "/api"
	http.HandleFunc("/api", handleRequest)

	// Start the server
	port := 8080
	fmt.Printf("Server running on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
