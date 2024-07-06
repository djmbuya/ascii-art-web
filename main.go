package main

import (
    "fmt"
    "html/template"
    "net/http"
)

// AsciiArt represents the ASCII art data
type AsciiArt struct {
    Text string
	Banner string
}

func main() {
    // Define the routes and their corresponding handlers
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/ascii", asciiHandler)
    http.HandleFunc("/submit", submitHandler)

    fmt.Println("Starting server on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

// indexHandler handles the root URL and renders the HTML template
func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Create a new AsciiArt instance
    art := &AsciiArt{}

    // Parse and execute the HTML template
    tmpl, _ := template.ParseFiles("index.html")
    tmpl.Execute(w, art)
}

// asciiHandler handles the "/ascii" URL and retrieves the ASCII art
func asciiHandler(w http.ResponseWriter, r *http.Request) {
    // Get the ASCII art text from the request
    text := r.FormValue("text")

    // Create a new AsciiArt instance with the retrieved text
    art := &AsciiArt{
        Text: text,
    }

    // Parse and execute the HTML template
    tmpl, _ := template.ParseFiles("index.html")
    tmpl.Execute(w, art)
}

// submitHandler handles the "/submit" URL and processes the form submission
func submitHandler(w http.ResponseWriter, r *http.Request) {
    // Get the ASCII art text from the form
    text := r.FormValue("text")
	banner := r.FormValue("banner")

    // Create a new AsciiArt instance with the submitted text
    art := &AsciiArt{
        Text: text,
    }
art = ascii.PrintAscii(text, banner)
    // Redirect the user to the "/ascii" URL with the ASCII art data
    http.Redirect(w, r, "/ascii?text="+text, http.StatusSeeOther)
}
