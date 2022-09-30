package main

import (
	"fmt"
	"strings"
)

func reverse(text string) string {
	result := ""
	for i := len(text) - 1; i >= 0; i-- {
		result += string(text[i])
	}

	return result
}

func isPalindrome(text string) bool {
	return text == reverse(text)
}

func main() {
	// Some basic string functions
	s1 := "   This is a super text   "
	fmt.Println("Lower text: [" + strings.ToLower(s1) + "]")
	fmt.Println("Upper text: [" + strings.ToUpper(s1) + "]")
	fmt.Println("Replacing first 's' with 'z': [" + strings.Replace(s1, "s", "z", 1) + "]")
	fmt.Println("Replacing all 's' with 'z': [" + strings.ReplaceAll(s1, "s", "z") + "]")
	fmt.Println("Supressing blank spaces: [" + strings.Trim(s1, " ") + "]")

	// Reverse and palindrome
	s2 := "mouse"
	fmt.Println("Reverse: " + reverse(s2))
	fmt.Println("Is '", s2, "' a palindrome?: ", isPalindrome(s2))

	s3 := "rotator"
	fmt.Println("Reverse: " + reverse(s3))
	fmt.Println("Is '", s3, "' a palindrome?: ", isPalindrome(s3))

	// Join an array of strings into a string separated by ;
	texts := []string{strings.Trim(s1, " "), s2, s3}
	s5 := strings.Join(texts, ";")
	fmt.Println("Joined strings: ", s5)

	// Split a string into an array, with texts separated by ;
	splitted := strings.Split(s5, ";")
	fmt.Println("Splitted strings:")

	for _, s := range splitted {
		fmt.Println("-", s)
	}

}
