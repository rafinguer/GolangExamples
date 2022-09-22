package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("--- Read from memory ---")
	basicCSVExample()
	fmt.Println("--- Read from file row by row ---")
	readCSVFile("contacts.csv")
	fmt.Println("--- Read from file all rows ---")
	readAllCSVFile("contacts.csv")
	fmt.Println("--- Writing data to a CSV file ---")
	writeCSVFile("authors.csv")
}

// Basic example for CSV
// CSV format in a string
func basicCSVExample() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemet","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("bad column:", pe.Column)
				fmt.Println("bad line:", pe.Line)
				fmt.Println("Error reported:", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

// This function reads the CSV file row by row
func readCSVFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#'
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

// This function reads the complete CSV file (all rows)
func readAllCSVFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#'
	r.Comma = ';'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)

	for i := 0; i < len(records); i++ {
		fmt.Printf("#%d-%s\n", i, records[i])
		fmt.Printf("   %s-%s$\n", records[i][1], records[i][3])
	}

}

// This function writes data to a CSV file
func writeCSVFile(filename string) {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemet", "gri"},
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	//w.WriteAll(data)
	w.Comma = ';'

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	err = w.Error()

	if err != nil {
		log.Fatal(err)
	}
}
