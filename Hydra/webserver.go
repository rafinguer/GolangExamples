package main

import (
	"Hydra/hlogger"
	"fmt"
	"net/http"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra Web Service...")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprintf(w, "Welcome to the Hydra Webserver")
	logger.Println("Received an http Get request on root url")
}
