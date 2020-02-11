package tasks

import (
	users "taskworks/app/models/users"
)

type Task struct {
	Title      string
	Owner      users.User
	Observer   int
	Department int
	Status     int
	User       users.User
}

func SaveTask(task *Task) *Task {
	return task
}
