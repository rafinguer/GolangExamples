package main

import (
	"fmt"
	"os"
)

// This function gives the arguments passed to the application
// from the command line. Example:
// $ go run osargs.go arg1 arg2 lastarg
func main() {
	fmt.Printf("Command: %#v\n", os.Args)
	fmt.Printf("Elements: %d\n", len(os.Args))
	fmt.Printf("Path: %s\n", os.Args[0])
	fmt.Println("Arguments:")

	for idx, arg := range os.Args {
		if idx > 0 {
			fmt.Println("- ", arg)
		}
	}
}
