package main

import "fmt"

func main() {
	fmt.Println("--- for min to max incrementing 1 ---")
	// Classical for: from min to max values incrementing iterator
	for i := 1; i <= 10; i++ {
		fmt.Println("Number: ", i)
	}

	// While condition
	fmt.Println("--- while condition ---")

	var i int = 1

	for i <= 10 {
		fmt.Println("Number: ", i)
		i++
	}

	// While condition
	fmt.Println("--- while condition (using continue) ---")

	i = 0

	for i <= 10 {
		i++

		// If noun number skip
		if i%2 != 0 {
			continue
		}

		// Pair number
		fmt.Println("Number: ", i)
	}

	// While true (infinite loop)
	// You must to use break sentence to break
	fmt.Println("--- while true (using break) ---")

	i = 1

	for {
		if i > 10 {
			break
		}

		fmt.Println("Number: ", i)
		i++
	}

	// For each/in
	fmt.Println("--- For each/in arrays ---")

	names := []string{"Rafael", "Edu", "Nerea", "Clemen", "Loli"}

	for idx, name := range names {
		fmt.Println("#", idx, ": ", name)
	}

	fmt.Println("--- For each/in with maps ---")

	var ColorMap = make(map[string]string)
	ColorMap["White"] = "#FFFFFF"
	ColorMap["Black"] = "#000000"
	ColorMap["Red"] = "#FF0000"
	ColorMap["Green"] = "#00FF00"
	ColorMap["Blue"] = "#0000FF"

	for colorName, colorValue := range ColorMap {
		fmt.Println("Color '", colorName, "': ", colorValue)
	}

}
