package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to represent the JSON response
type Message struct {
	Text string `json:"message"`
}

func main() {
	// Defining the "/hello" route to respond to the GET method
	http.HandleFunc("/hello", helloHandler)

	// Initializing the server on port 8080
	fmt.Println("Server started at http://localhost:8080/hello")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}

// Handler for the "/hello" route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Creating an instance of the Message struct
	message := Message{Text: "Hello, world!"}

	// Converting the struct to JSON
	jsonResponse, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Setting the response header with content type as application/json
	w.Header().Set("Content-Type", "application/json")

	// Writing the JSON response to the response body
	w.Write(jsonResponse)
}
