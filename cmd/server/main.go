package main

import (
	"log"      // Стандартный логгер
	"net/http" // HTTP сервер
	"person_service/pkg/logging"
	"person_service/pkg/storage"
	"person_service/pkg/transport/https"
)

func main() {
	// Инициализация логгера
	logger := logging.NewLogger()
	logger.LogInfo("Starting application...")

	// Инициализация хранилища
	store := storage.NewStorage(logger)
	logger.LogInfo("Storage initialized")

	// Регистрация обработчиков
	http.HandleFunc("/save", https.SaveHandler(store, logger))
	http.HandleFunc("/getPerson", https.GetPersonHandler(store, logger))
	logger.LogInfo("Handlers registered")

	// Запуск сервера
	logger.LogInfo("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.LogError("Server failed: " + err.Error())
		log.Fatal(err) // Аварийное завершение
	}
}
