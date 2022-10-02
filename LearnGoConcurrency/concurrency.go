package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServers(uri string, channel chan string) {
	_, err := http.Get(uri)

	if err != nil {
		channel <- uri + " isn't working: " + err.Error()
	} else {
		channel <- uri + " is working: "
	}
}

func main() {
	start := time.Now()

	channel := make(chan string)

	servers := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.youtube.com",
		"https://es.yahoo.com/",
		"https://duckduckgo.com/",
		"https://www.startpage.com/",
		"https://www.nytimes.com/international/",
		"https://www.bbc.com/mundo",
	}

	for _, server := range servers {
		go checkServers(server, channel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	totalTime := time.Since(start)

	fmt.Println("Total time: ", totalTime)
}
