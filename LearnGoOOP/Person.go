package main

import "fmt"

// Defining the Person's properties
type Person struct {
	name string
	age  int
}

// Constructor
func (p *Person) Constructor(name string, age int) {
	p.name = name
	p.age = age
}

// Setting the name of the Person
func (p *Person) SetName(name string) {
	p.name = name
}

// Setting the age of the Person
func (p *Person) SetAge(age int) {
	p.age = age
}

// Getting the name of the Person
func (p *Person) GetName() string {
	return p.name
}

// Getting the age of the Person
func (p *Person) GetAge() int {
	return p.age
}

// Function that returns the Person's properties
func (p *Person) String() string {
	return fmt.Sprintf("name: %s, age: %d", p.name, p.age)
}
