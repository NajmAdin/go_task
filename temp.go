package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// API routes

	// Serve files from static folder
	http.Handle("/", http.FileServer(http.Dir("./source")))

	// Serve api /hi
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))

}
