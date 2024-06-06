package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
	CSSFile string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/navbar.html",
		"templates/sidemenu.html",
		"templates/content.html",
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

	http.HandleFunc("/", handler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
