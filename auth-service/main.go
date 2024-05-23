package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached the Auth Service")
	})

	fmt.Println("Auth Service listening on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
