// main.go
// Этот модуль является точкой входа для сервера приложения Workmate-GO.
// Он загружает конфигурацию, настраивает маршрутизатор и запускает HTTP-сервер.

package main

import (
	"log"
	"net/http"

	"github.com/ekTeZy/Workmate-GO/internal/config"
	"github.com/ekTeZy/Workmate-GO/internal/handler"
)

func main() {
	// Загрузка конфигурации приложения
	cfg := config.LoadConfig()

	// Создание нового маршрутизатора
	mux := http.NewServeMux()

	// Настройка маршрутов
	mux.HandleFunc("/", rootHandler)            // Корневой маршрут
	mux.HandleFunc("/task/", handler.GetTask)   // Маршрут для получения задачи
	mux.HandleFunc("/task", handler.CreateTask) // Маршрут для создания задачи

	// Формирование адреса сервера
	addr := ":" + cfg.Port

	// Логирование запуска сервера
	log.Println("Сервер запущен на", addr)

	// Запуск HTTP-сервера
	if err := http.ListenAndServe(addr, mux); err != nil {
		// Логирование ошибки при запуске сервера
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

// Обработчик корневого маршрута
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Возврат статуса 200 OK и приветственного сообщения
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Добро пожаловать!"))
}
