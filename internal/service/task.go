package service

import (
	"log"
	"math/rand"
	"time"

	"github.com/ekTeZy/Workmate-GO/internal/model"
	"github.com/ekTeZy/Workmate-GO/internal/repository"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateTask() *model.Task {
	id := uuid.NewString()

	task := &model.Task{
		ID:     id,
		Status: model.StatusPending,
	}

	log.Printf("[SERVICE][NEW] Создана задача: ID=%s", id)

	return task
}

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

func GetTaskByID(id string) (*model.Task, bool) {
	return repository.GetTaskByID(id)
}
