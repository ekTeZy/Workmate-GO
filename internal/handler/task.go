package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ekTeZy/Workmate-GO/internal/model"
	"github.com/ekTeZy/Workmate-GO/internal/service"
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
	w.Header().Set("Content-Type", "application/json")

	task, err := service.StartTask()
	if err != nil {
		log.Printf("[ERROR] Не удалось создать задачу: %v", err)
		http.Error(w, "Что-то пошло не так с задачей", http.StatusInternalServerError)
		return
	}

	log.Printf("[HANDLER] Создана задача с ID %s", task.ID)

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/task/")
	if id == "" {
		log.Printf("[WARN] Пустой ID в запросе")
		http.Error(w, "ID задачи не указан", http.StatusBadRequest)
		return
	}

	task, found := service.GetTaskByID(id)
	if !found {
		log.Printf("[WARN] Задача с ID %s не найдена", id)
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	log.Printf("[HANDLER] Получена задача ID %s, статус: %s", task.ID, task.Status)

	data := TaskResponse{
		ID:     task.ID,
		Status: string(task.Status),
		Result: task.Result,
		Error:  task.Error,
	}

	resp := Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Data:       data,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
