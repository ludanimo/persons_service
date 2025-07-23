package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"person_service/pkg/logging"
	"person_service/pkg/storage"
	httptransport "person_service/pkg/transport/https"
	"strings"
	"testing"
)

// MockLogger должен точно повторять структуру реального Logger
type MockLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewMockLogger() *logging.Logger {
	return &logging.Logger{
		InfoLogger:  log.New(&nullWriter{}, "INFO: ", 0),
		ErrorLogger: log.New(&nullWriter{}, "ERROR: ", 0),
	}
}

// nullWriter для подавления вывода логов в тестах
type nullWriter struct{}

func (w *nullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *MockLogger) LogInfo(message string) {
	m.infoLogger.Println(message)
}

func (m *MockLogger) LogError(message string) {
	m.errorLogger.Println(message)
}

func TestSaveHandler(t *testing.T) {
	// Инициализация
	logger := NewMockLogger()
	store := storage.NewStorage(logger) // Используем NewStorage как в вашем коде
	handler := httptransport.SaveHandler(store, logger)

	// Тестовый запрос
	req, _ := http.NewRequest("POST", "/save?ID=1&Name=John", nil)
	rr := httptest.NewRecorder()

	// Вызов обработчика
	handler.ServeHTTP(rr, req)

	// Проверки
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}

	// Проверка сохранения в хранилище
	name, ok := store.Get(1)
	if !ok || name != "John" {
		t.Errorf("Expected John in storage, got %v", name)
	}
}

func TestGetPersonHandler(t *testing.T) {
	// Инициализация
	logger := NewMockLogger()
	store := storage.NewStorage(logger)
	store.Save(1, "John") // Используем напрямую метод хранилища

	handler := httptransport.GetPersonHandler(store, logger)

	// Тестовый запрос
	req, _ := http.NewRequest("GET", "/getPerson?ID=1", nil)
	rr := httptest.NewRecorder()

	// Вызов обработчика
	handler.ServeHTTP(rr, req)

	// Проверки
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}

	expected := `{"name":"John"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Expected %s, got %s", expected, rr.Body.String())
	}
}
