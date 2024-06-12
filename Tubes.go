package main

import (
	"Tubes/messages"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// PageData represents the data structure for the HTML templates
type PageData struct {
	Title   string
	Message string
	CSSFile string
}

// HandleIncomingMessage handles incoming messages
func HandleIncomingMessage(w http.ResponseWriter, r *http.Request) {
	var msg messages.Message

	// Parse the incoming JSON message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid message format", http.StatusBadRequest)
		return
	}

	// Handle the message based on its type
	response, err := messages.HandleMessage(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ServeTemplate serves the HTML templates
func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/navbar.html",
		"templates/sidemenu.html",
		"templates/content.html",
		"templates/connections.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check for a query parameter to decide which CSS file to use
	cssFile := "static/style.css"
	if r.URL.Query().Get("alt") == "true" {
		cssFile = "static/alternative.css"
	}

	data := PageData{
		Title:   "Welcome to My Website",
		Message: "Hello, World!",
		CSSFile: cssFile,
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle incoming messages
	http.HandleFunc("/message", HandleIncomingMessage)

	// Serve HTML templates
	http.HandleFunc("/", ServeTemplate)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
