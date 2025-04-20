package repository

import (
	"fmt"
	"sync"

	"github.com/ekTeZy/Workmate-GO/internal/model"
)

var (
	tasks = make(map[string]*model.Task)
	mu    sync.RWMutex
)

// SaveTask сохраняет задачу в памяти.
// Если задача с таким ID уже существует, возвращает ошибку.
func SaveTask(task *model.Task) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := tasks[task.ID]; ok {
		return fmt.Errorf("задача с id %s уже существует", task.ID)
	}

	tasks[task.ID] = task

	return nil
}

// UpdateStatus обновляет статус задачи по ID.
// Если задача с таким ID не существует, возвращает ошибку.
// Если статус задачи изменен на "Выполнено" или "Ошибка", обновляет результат задачи.
func UpdateStatus(id string, status model.TaskStatus, taskResult string) error {
	mu.Lock()
	defer mu.Unlock()

	task, ok := tasks[id]
	if !ok {
		return fmt.Errorf("задача с id %s не существует", id)
	}

	task.Status = status

	if status == model.StatusDone || status == model.StatusError {
		task.Result = taskResult
	}

	return nil
}

// GetTaskByID возвращает задачу по ID.
// Если задача с таким ID не существует, возвращает false.
func GetTaskByID(id string) (*model.Task, bool) {
	mu.RLock()
	defer mu.RUnlock()

	task, ok := tasks[id]
	return task, ok
}
