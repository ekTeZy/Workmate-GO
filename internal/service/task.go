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

	log.Printf("[NEW] Создана задача с ID %s", id)

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

func RunTask(t *model.Task) {
	log.Printf("[RUNNING] Задача %s: статус установлен в running", t.ID)
	repository.UpdateStatus(t.ID, model.StatusRunning, "")

	time.Sleep(5 * time.Second)

	taskResult := rand.Intn(2) == 1
	if !taskResult {
		repository.UpdateStatus(t.ID, model.StatusError, "Нет данных")
		log.Printf("[ERROR] Задача %s завершилась с ошибкой: результат = \"Нет данных\"", t.ID)
	} else {
		repository.UpdateStatus(t.ID, model.StatusDone, "Результат готов")
		log.Printf("[DONE] Задача %s завершена: результат = \"Результат готов\"", t.ID)
	}
}

func GetTaskByID(id string) (*model.Task, bool) {
	return repository.GetTaskByID(id)
}
