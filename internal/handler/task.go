package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ekTeZy/Workmate-GO/internal/model"
	"github.com/ekTeZy/Workmate-GO/internal/service"
	"github.com/google/uuid"
)

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type TaskResponse struct {
	ID     string `json:"task_id"`
	Status string `json:"task_status"`
	Result string `json:"task_result,omitempty"`
	Error  string `json:"task_error,omitempty"`
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("[HANDLER][ERROR] Метод %s не поддерживается", r.Method)
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	task, err := service.StartTask()
	if err != nil {
		log.Printf("[HANDLER][ERROR] Ошибка при запуске задачи: %v", err)
		http.Error(w, "Что-то пошло не так с задачей", http.StatusInternalServerError)
		return
	}

	log.Printf("[HANDLER][POST] Задача создана: ID=%s", task.ID)

	data := TaskResponse{
		ID:     task.ID,
		Status: string(task.Status),
		Result: task.Result,
		Error:  task.Error,
	}

	respStatus := "success"
	if task.Status == model.StatusError {
		respStatus = "failed"
	}

	response := Response{
		Status:     respStatus,
		StatusCode: http.StatusCreated,
		Data:       data,
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("[HANDLER][ERROR] Ошибка сериализации ответа: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/task/")
	if id == "" {
		log.Println("[HANDLER][WARN] Пустой ID задачи в запросе")
		http.Error(w, "ID задачи не указан", http.StatusBadRequest)
		return
	}

	if _, err := uuid.Parse(id); err != nil {
		log.Printf("[HANDLER][WARN] Неверный формат ID: %s", id)
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	task, found := service.GetTaskByID(id)
	if !found {
		log.Printf("[HANDLER][WARN] Задача не найдена: ID=%s", id)
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	log.Printf("[HANDLER][GET] Задача найдена: ID=%s, статус=%s", task.ID, task.Status)

	data := TaskResponse{
		ID:     task.ID,
		Status: string(task.Status),
		Result: task.Result,
		Error:  task.Error,
	}

	response := Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Data:       data,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("[HANDLER][ERROR] Ошибка сериализации ответа: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
