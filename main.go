package main

import (
	// a "ascii/ascii_art"
	"fmt"
	"log"
	"net/http"

	//"html/template"
	"ascii/handlers"
)

// main starts the HTTP server.
func main() {
	fmt.Println("Server is starting...")
	http.HandleFunc("/", handlers.Handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Server up at port 8080\nhttp status :", http.StatusOK)
	// Openbrowser("http.localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
