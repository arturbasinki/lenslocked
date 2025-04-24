package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome website!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Server is running on http://localhost:3030")
	if err := http.ListenAndServe(":3030", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
