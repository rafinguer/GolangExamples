package main

import (
	"fmt"
	"log"
	"net/http"
)

// Note: If you want the server is automatically self reinitilizes
// install the "fresh" framework:
// $ go get github.com/pilu/fresh
// Go to the app directory and execute:
// $ fresh

// Handlers
func Hello(rw http.ResponseWriter, r *http.Request) {
	// Return content to client
	fmt.Fprintln(rw, "Hello, Go World!!!")

	// Printing Request info
	fmt.Println("Method: ", r.Method)
	//fmt.Println("Header	: ", r.Header)
	//fmt.Println("Body: ", r.Body)
}

// http://localhost:3000?name=<name>&age=<age>
func Greet(rw http.ResponseWriter, r *http.Request) {
	// Return content to client
	//fmt.Println("URL: ", r.URL)  // Get the URL
	//fmt.Println("URL: ", r.URL.RawQuery)  // Separate arguments from the URL
	//fmt.Println("URL: ", r.URL.Query())  // Get the arguments into a map
	name := r.URL.Query().Get("name") // Get argument "name" (key of the map) (string)
	age := r.URL.Query().Get("age")   // Get argument "name" (key of the map) (string)
	fmt.Fprintln(rw, "Hello, "+name+"!!!"+". You are "+age+" years old")
}

func ErrorPage(rw http.ResponseWriter, r *http.Request) {
	// Return an error
	http.Error(rw, "This a fucking error!!!", http.StatusTeapot)
}

// Main function
func main() {
	// Router
	//http.HandleFunc("/", Hello)
	//http.HandleFunc("/error", ErrorPage)
	//http.HandleFunc("/greet", Greet)

	// Server with router
	//fmt.Println("Running web server on http://localhost:3000")
	//log.Fatal(http.ListenAndServe(":3000", nil))  // with http handlers

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/error", ErrorPage)
	mux.HandleFunc("/greet", Greet)

	// Server with mux
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	fmt.Println("Running web server on http://localhost:3000")
	log.Fatal(server.ListenAndServe()) // with mux handlers
}
