package handler

import (
	"encoding/json"
	"net/http"

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
		http.Error(w, "Что-то пошло не так с задачей", http.StatusInternalServerError)
		return
	}

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
