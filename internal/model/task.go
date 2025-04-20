package model

type TaskStatus string

const (
	StatusPending TaskStatus = "pending"
	StatusRunning TaskStatus = "running"
	StatusDone    TaskStatus = "done"
	StatusError   TaskStatus = "error"
)

type Task struct {
	ID     string
	Status TaskStatus
	Result string
	Error  string
}
