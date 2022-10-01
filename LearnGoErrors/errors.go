package main

import (
	"errors"
	"fmt"
	"os"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}

	return a / b, nil
}

func main() {
	result, err := divide(12, 6)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("12 / 6 = ", result)

	result, err = divide(3, 0)

	if err != nil {
		fmt.Println("Error found:", err)
	}

	// Recover controls the panic executions, avoiding the application break
	// Defer func is executed at the end of the application execution
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Oops! It seems like the application was finished in an incorrect way, and it was caused by a panic")
		}
	}()

	// Managing a file
	if file, error := os.Open("example.txt"); error != nil {
		//fmt.Println("Cannot open file:", error)

		// panic() breaks abruptaly the application due to a critical error
		panic(fmt.Sprintf("Cannot open file: %s", error))
	} else {
		defer func() { // Defer is executed at the end (after all file operations)
			fmt.Println("Closing file")
			file.Close()
		}()

		content := make([]byte, 255)
		long, _ := file.Read(content)
		contentFile := string(content[:long])
		fmt.Println("Content: [" + contentFile + "]")
	}

}
