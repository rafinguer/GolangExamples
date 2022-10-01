package main

import "fmt"

// Person <- Employee (Employee inherits from Person)

// Defining the Employee properties
type Employee struct {
	Person // Inheritance from the class Person
	salary float64
}

// Setting the employee's salary
func (e *Employee) SetSalary(s float64) {
	e.salary = s
}

// Getting the employee's salary
func (e *Employee) GetSalary() float64 {
	return e.salary
}

// Overriging the String() function
func (e *Employee) String() string {
	return fmt.Sprintf("%s, salary: %.2f", e.Person.String(), e.salary)
}
