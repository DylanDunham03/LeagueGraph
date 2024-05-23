package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached the gRPC Gateway")
	})

	fmt.Println("gRPC Gateway listening on http://localhost:8082")
	http.ListenAndServe(":8082", nil)
}
