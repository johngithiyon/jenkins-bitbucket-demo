package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Jenkins + Bitbucket!")
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
