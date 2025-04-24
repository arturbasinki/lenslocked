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

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h1>404 Not Found</h1><p>The page you are looking for does not exist.</p>")
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	// Start the server on port 3000
	fmt.Println("Server is running on http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
