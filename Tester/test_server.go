package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Respond with "Hello" for GET requests
		fmt.Fprintln(w, "Hello")
	} else {
		// Respond with a 405 Method Not Allowed for other request types
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Set up the route and handler
	http.HandleFunc("/", helloHandler)

	// Start the server on localhost:3131
	fmt.Println("Server running on http://localhost:3131")
	if err := http.ListenAndServe(":3131", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
