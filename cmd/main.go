package main

import (
	"log/slog"
	"net/http"
	"persons_service/internal/interface/controller"
	"persons_service/internal/interface/repository"
	"persons_service/internal/usecase/person"
)

func main() {

	slog.Info("Starting application...")

	// Инициализация репозитория
	repo := repository.NewInMemoryRepository()
	slog.Info("Initializing repository...")

	// Инициализация usecase
	usecase := person.NewPersonUsecase(repo)

	// Инициализация контроллера
	handler := controller.NewPersonHandler(usecase)

	// Регистрация обработчиков
	http.HandleFunc("/save", handler.SaveHandler())
	http.HandleFunc("/getPerson", handler.GetHandler())

	// Запуск сервера
	slog.Info("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Server failed: " + err.Error())
		panic(err)
	}
}
