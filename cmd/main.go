package main

import (
	"log/slog"
	"net/http"
	"persons_service/internal/entity"
	"persons_service/internal/repository"
	"persons_service/internal/service"
)

func main() {
	slog.Info("Starting application...")

	// Инициализация репозитория
	repo := repository.NewInMemoryRepository()
	slog.Info("Initializing repository...")

	// Инициализация бизнес сценариев
	validationService := entity.NewPersonService(repo)

	// Инициализация обработчика
	handler := service.NewPersonHandler(validationService)

	// Регистрация обработчиков
	http.HandleFunc("/person", handler.CreateHandler())
	http.HandleFunc("/id", handler.GetHandler())

	// Запуск сервера
	slog.Info("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Server failed: " + err.Error())
		panic(err)
	}
}
