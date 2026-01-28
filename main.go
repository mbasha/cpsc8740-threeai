package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"threeai/calculator"
	"threeai/tictactoe"
	"threeai/todolist"
)

var (
	templateDir = "templates"
	staticDir   = "static"
)

func main() {
	// Ensure data directories exist
	os.MkdirAll("data", 0755)

	// Parse templates
	homeTemplate := template.Must(template.ParseFiles(filepath.Join(templateDir, "index.html")))

	// File server for static assets
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Home route - app selection
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		homeTemplate.Execute(w, nil)
	})

	// Calculator routes
	calculator.RegisterRoutes(http.DefaultServeMux)

	// Tic Tac Toe routes
	tictactoe.RegisterRoutes(http.DefaultServeMux)

	// To-Do List routes
	todolist.RegisterRoutes(http.DefaultServeMux)

	port := ":8080"
	fmt.Printf("Starting server on http://localhost%s\n", port)
	fmt.Println("Navigate to http://localhost:8080 to get started")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
