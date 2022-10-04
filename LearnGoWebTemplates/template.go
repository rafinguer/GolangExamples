package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// Handlers
func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola")
	// Getting current directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Make the correct path
	if strings.Contains(wd, "/") { // Linux
		wd += "/index.html"
	} else { // Windows
		wd += "\\index.html"
	}

	fmt.Println("Path: ", wd)

	t, error := template.ParseFiles(wd)

	//t, error := template.ParseFiles("index.html")

	if err != nil {
		panic(error)
	} else {
		t.ExecuteTemplate(w, wd, nil)
	}
}

func main() {
	// mux definitions
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// Server with mux
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	fmt.Println("Running web server on http://localhost:3000")
	log.Fatal(server.ListenAndServe()) // with mux handlers
}
