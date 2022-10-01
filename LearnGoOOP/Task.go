package main

import "fmt"

type Task struct {
	name      string
	completed bool
}

// Marking the task as completed
func (t *Task) markCompleted() {
	t.completed = true
}

// Showing the task data
func (t *Task) String() string {
	return fmt.Sprintf("task: '%s', completed: %t", t.name, t.completed)
}
