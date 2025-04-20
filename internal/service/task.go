package service

import (
	"log"
	"math/rand"
	"time"

	"github.com/ekTeZy/Workmate-GO/internal/model"
	"github.com/ekTeZy/Workmate-GO/internal/repository"
	"github.com/google/uuid"
)

func CreateTask() *model.Task {
	id := uuid.NewString()

	task := &model.Task{
		ID:     id,
		Status: model.StatusPending,
	}

	return task
}

func StartTask() (*model.Task, error) {
	task := CreateTask()
	err := repository.SaveTask(task)
	if err != nil {
		return task, err
	}

	go RunTask(task)

	return task, nil
}

func RunTask(t *model.Task) (*model.Task, error) {
	log.Printf("Задача %s запущена", t.ID)

	repository.UpdateStatus(t.ID, model.StatusRunning, "")

	time.Sleep(5 * time.Second)

	rand.Seed(time.Now().UnixNano())
	taskResult := rand.Intn(2) == 1
	if !taskResult {
		repository.UpdateStatus(t.ID, model.StatusError, "Нет данных")
		log.Printf("Задача %s завершена с ошибкой", t.ID)
	} else {
		repository.UpdateStatus(t.ID, model.StatusDone, "Результат готов")
		log.Printf("Задача %s завершена успешно", t.ID)
	}

	return t, nil
}

func GetTaskByID(id string) (*model.Task, bool) {
	return repository.GetTaskByID(id)
}
