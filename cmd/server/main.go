package main

import (
	"log"
	"net/http"

	"github.com/ekTeZy/Workmate-GO/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	addr := ":" + cfg.Port
	log.Println("Сервер запущен на", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Добро пожаловать!"))
}
