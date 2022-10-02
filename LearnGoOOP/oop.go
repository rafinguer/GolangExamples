package main

import (
	"fmt"
)

func main() {
	personExample()
	employeeExample()
	taskExample()
	taskListExample()
	relationOneToOneExample()
	relationOneToManyExample()
	interfaceExample()
}

// This function shows how to use the Person class
func personExample() {
	fmt.Println("\n---\nPerson Example\n---")
	//var person1 Person = Person{name: "Rafael", age: 50}
	var p Person = Person{"Rafael", 50} // In the same order you can omit property names
	fmt.Println("Person: ", p.String())
	p.SetName("Nerea")
	p.SetAge(20)
	fmt.Println("Name: ", p.GetName())
	fmt.Println("Age: ", p.GetAge())

	// Creating a new Person via Constructor
	p2 := Person{}
	p2.Constructor("Edu", 52)
	fmt.Println("Person: ", p2.String())
}

// This function shows how to use the Employee class
func employeeExample() {
	fmt.Println("\n---\nEmployee Example\n---")
	// Employee inherits from Person, adding the salary property and its methods
	// When create a new employee, only its own properties can be initialized
	var e Employee = Employee{salary: 1028.1234}
	fmt.Println("The salary of the employee is:", e.GetSalary(), "euros")
	e.SetName("Rafael")
	e.SetAge(50)
	e.SetSalary(1234.5678)       // changing the salary
	fmt.Println("Employee: ", e) // {{"Rafael", 50}, 1234.5678}  // {{<Person>} <Employee>}
	fmt.Println("Employee's String(): ", e.String())
	fmt.Println("Employee's name", e.Person.name) // Property inherited from Person
	fmt.Println("Employee's age", e.Person.age)   // Property inherited from Person
	fmt.Println("Employee's salary", e.salary)    // Property owned by Employee
}

// This function shows how to use the Task class
func taskExample() {
	fmt.Println("\n---\nTask Example\n---")
	var t Task = Task{"Task1", false}
	fmt.Println("Task before:", t.String())
	t.markCompleted()
	fmt.Println("Task after:", t.String())

	t2 := Task{
		name:      "Task #2",
		completed: false,
	}
	fmt.Println("Task2: ", t2)
}

// This function shows how to use the TaskList class
func taskListExample() {
	fmt.Println("\n---\nTaskList Example\n---")

	// Creating and initializing a task list
	myTaskList := TaskList{tasks: []*Task{
		{name: "Learn Go", completed: false},
		{name: "Holidays", completed: true},
		{name: "Buy Clemen's Birthday Gift", completed: false},
	}}

	for i, v := range myTaskList.tasks {
		fmt.Printf("Task #%d -> name: %s, completed: %t\n", i, v.name, v.completed)
	}

	fmt.Println("\n---")

	// Creating individual tasks
	var t1 Task = Task{"Buy milk", true}
	var t2 Task = Task{"Visit dentist", false}
	var t3 Task = Task{"Pay Go course", true}

	// Adding individual tasks to the TaskList
	myTaskList.appendTask(&t1)
	myTaskList.appendTask(&t2)
	myTaskList.appendTask(&t3)

	for i, v := range myTaskList.tasks {
		fmt.Printf("Task #%d -> name: %s, completed: %t\n", i, v.name, v.completed)
	}

	fmt.Println("\n---")

	// Removing task at index 1 (Holidays) and 3 (Buy milk) and 5 (Pay Go course)
	// Put attention to the deletion is made from the highest index to the lowest one
	// If you delete form the lowest, the list task will decrease and indexes won't match
	myTaskList.removeTask(5)
	myTaskList.removeTask(3)
	myTaskList.removeTask(1)

	for i, v := range myTaskList.tasks {
		fmt.Printf("Task #%d -> name: %s, completed: %t\n", i, v.name, v.completed)
	}
}

// This function demonstrates how to create a relation one to one between two classes
// For this example, we use User and Student classes
// In this case, Student inherits from User
// One Student only can be one user
func relationOneToOneExample() {
	fmt.Println("\n---\nRelation One to One Example\n---")
	rafael := User{"Rafael", "rafael@gmail.com", true}    // User object
	rafaelStudent := Student{user: rafael, studentId: 1}  // Student object
	fmt.Println("Student: ", rafaelStudent)               // {{<User>} <Student>}
	fmt.Println("Student ID:", rafaelStudent.studentId)   // Student property
	fmt.Println("Student name:", rafaelStudent.user.name) // User object inside Student object
	fmt.Println("Student email:", rafaelStudent.user.email)
	fmt.Println("Student active:", rafaelStudent.user.active)
}

// This function demonstrates how to create a relation one to many between two classes
// For this example, we use Student and Course classes
// In this case, one Student can learn many courses
func relationOneToManyExample() {
	fmt.Println("\n---\nRelation One to Many Example\n---")
	rafael := User{"Rafael", "rafael@gmail.com", true}   // User object
	rafaelStudent := Student{user: rafael, studentId: 1} // Student object

	c1 := Course{1, "Golang"}
	c2 := Course{2, "Groovy"}
	c3 := Course{3, "Rust"}
	rafaelStudent.courses = append(rafaelStudent.courses, c1, c2, c3)

	fmt.Println("Student: ", rafaelStudent)               // {{<User>} <Student>}
	fmt.Println("Student ID:", rafaelStudent.studentId)   // Student property
	fmt.Println("Student name:", rafaelStudent.user.name) // User object inside Student object
	fmt.Println("Student email:", rafaelStudent.user.email)
	fmt.Println("Student active:", rafaelStudent.user.active)

	for i, v := range rafaelStudent.courses {
		fmt.Printf("- Course #%d -> courseId: %d, title: %s\n", i, v.courseId, v.title)
	}
}

func interfaceExample() {
	fmt.Println("\n---\nRelation One to Many Example\n---")
	square := Square{24.0}
	rectangle := Rectangle{24.0, 15.5}
	circle := Circle{15.8}

	fmt.Println("Area of square: ", calculateShapeArea(&square))
	fmt.Println("Area of rectangle: ", calculateShapeArea(&rectangle))
	fmt.Println("Area of circle: ", calculateShapeArea(&circle))
	fmt.Println("Perimeter of square: ", calculateShapePerimeter(&square))
	fmt.Println("Perimeter of Rectangle: ", calculateShapePerimeter(&rectangle))
	fmt.Println("Perimeter of circle: ", calculateShapePerimeter(&circle))
}
