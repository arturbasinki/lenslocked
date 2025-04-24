package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to plain text or HTML
	// w.Header().Set("Content-Type", "text/plain")
	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Write a simple HTML response
	fmt.Fprint(w, "<h1>Welcome to my awesome website!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Write a simple HTML response
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">artur.basinski@dupa.com</a>.</p>")
}

type Router struct{}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "The page you are looking for does not exist.", http.StatusNotFound)
	}
}

func main() {
	// Create a new router
	router := &Router{}
	// Start the server on port 3000
	fmt.Println("Server is running on http://localhost:3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
