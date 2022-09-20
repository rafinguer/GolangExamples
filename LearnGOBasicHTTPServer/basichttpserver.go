package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", welcomePage)
	http.ListenAndServe(":8080", nil)
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to this Site</h1>")
}
