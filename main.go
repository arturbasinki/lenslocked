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
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	// Start the server on port 3030
	fmt.Println("Server is running on http://localhost:3030")
	if err := http.ListenAndServe(":3030", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
