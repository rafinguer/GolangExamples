package main

// This class manages a list of Tasks
type TaskList struct {
	tasks []*Task
}

// Adding new task to the list
func (tl *TaskList) appendTask(task *Task) {
	tl.tasks = append(tl.tasks, task)
}

// Deleting a given task from the list
func (tl *TaskList) removeTask(idx int) {
	tl.tasks = append(tl.tasks[:idx], tl.tasks[idx+1:]...)
}
