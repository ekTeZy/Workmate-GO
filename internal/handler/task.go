// Package handler содержит обработчики HTTP-запросов для управления задачами.
//
// Этот модуль включает в себя функции для создания и получения задач.
// Обработчики используют пакет service для взаимодействия с бизнес-логикой приложения.
//
// Доступные функции:
// - CreateTask: Создает новую задачу и возвращает её ID и статус.
// - GetTask: Получает информацию о задаче по её ID.
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

// Response представляет собой структуру ответа, содержащую статус, код состояния и данные.
type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

// TaskResponse представляет собой структуру ответа, содержащую информацию о задаче.
type TaskResponse struct {
	ID     string `json:"task_id"`
	Status string `json:"task_status"`
	Result string `json:"task_result,omitempty"`
	Error  string `json:"task_error,omitempty"`
}

// CreateTask обрабатывает HTTP-запросы на создание новой задачи.
// Если метод запроса не POST, возвращает ошибку 405.
// Создает задачу через service.StartTask и возвращает её ID и статус.
// В случае ошибки возвращает ошибку 500.
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

// GetTask обрабатывает HTTP-запросы на получение информации о задаче по её ID.
// Если ID не указан или имеет неверный формат, возвращает ошибку 400.
// Если задача не найдена, возвращает ошибку 404.
// В случае успеха возвращает информацию о задаче.
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
