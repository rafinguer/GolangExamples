package main

import "fmt"

type Course struct {
	courseId int
	title    string
}

func (c *Course) String() string {
	return fmt.Sprintf("Course -> courseId: %d, title: %s", c.courseId, c.title)
}
