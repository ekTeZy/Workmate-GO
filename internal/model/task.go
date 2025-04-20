package model

// TaskStatus представляет собой тип строки, используемый для обозначения статуса задачи.
type TaskStatus string

const (
	// StatusPending обозначает, что задача находится в состоянии ожидания.
	StatusPending TaskStatus = "pending"
	// StatusRunning обозначает, что задача выполняется.
	StatusRunning TaskStatus = "running"
	// StatusDone обозначает, что задача завершена успешно.
	StatusDone TaskStatus = "done"
	// StatusError обозначает, что задача завершена с ошибкой.
	StatusError TaskStatus = "error"
)

// Task представляет собой структуру, описывающую задачу.
type Task struct {
	// ID уникальный идентификатор задачи.
	ID string
	// Status текущий статус задачи.
	Status TaskStatus
	// Result результат выполнения задачи.
	Result string
	// Error сообщение об ошибке, если задача завершена с ошибкой.
	Error string
}
