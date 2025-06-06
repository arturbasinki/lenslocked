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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
  </li>
</ul>
`)
}

type Router struct{}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
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
