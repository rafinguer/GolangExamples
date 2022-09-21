package main

import "Hydra/hydrachat"

// Open a terminal in this folder and
//    go run hchatservertest.go
// After this, open some terminals executing hchatclient
func main() {
	hydrachat.Run(":2300")
}
