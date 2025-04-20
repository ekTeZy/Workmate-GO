// Package service содержит функции для управления задачами в системе.
package service

import (
	"log"
	"math/rand"
	"time"

	"github.com/ekTeZy/Workmate-GO/internal/model"
	"github.com/ekTeZy/Workmate-GO/internal/repository"
	"github.com/google/uuid"
)

// init инициализирует генератор случайных чисел с использованием текущего времени в наносекундах.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// CreateTask создает новую задачу с уникальным идентификатором и статусом Pending.
func CreateTask() *model.Task {
	id := uuid.NewString()

	task := &model.Task{
		ID:     id,
		Status: model.StatusPending,
	}

	log.Printf("[SERVICE][NEW] Создана задача: ID=%s", id)

	return task
}

// StartTask создает новую задачу и сохраняет её в репозитории.
// Если сохранение прошло успешно, задача запускается в фоновом режиме.
func StartTask() (*model.Task, error) {
	task := CreateTask()

	err := repository.SaveTask(task)
	if err != nil {
		log.Printf("[SERVICE][ERROR] Не удалось сохранить задачу ID=%s: %v", task.ID, err)
		return task, err
	}

	log.Printf("[SERVICE][START] Задача сохранена и запущена в фоне: ID=%s", task.ID)
	go RunTask(task)

	return task, nil
}

// RunTask выполняет задачу, имитируя долгую операцию.
// В зависимости от результата выполнения, задача обновляется со статусом Done или Error.
func RunTask(t *model.Task) {
	log.Printf("[SERVICE][RUN] Задача ID=%s: установка статуса running", t.ID)
	repository.UpdateStatus(t.ID, model.StatusRunning, "")

	time.Sleep(5 * time.Second) // имитация долгой задачи

	taskResult := rand.Intn(2) == 1
	if !taskResult {
		repository.UpdateStatus(t.ID, model.StatusError, "Нет данных")
		log.Printf("[SERVICE][FAIL] Задача ID=%s завершена с ошибкой: результат=\"Нет данных\"", t.ID)
	} else {
		repository.UpdateStatus(t.ID, model.StatusDone, "Результат готов")
		log.Printf("[SERVICE][DONE] Задача ID=%s завершена успешно: результат=\"Результат готов\"", t.ID)
	}
}

// GetTaskByID получает задачу по её идентификатору из репозитория.
func GetTaskByID(id string) (*model.Task, bool) {
	return repository.GetTaskByID(id)
}
