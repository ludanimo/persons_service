package logging

import (
	"log"
	"os"
)

// Logger структура для логгирования
type Logger struct {
	InfoLogger  *log.Logger // Логгер информационных сообщений
	ErrorLogger *log.Logger // Логгер ошибок
}

// NewLogger создает новый экземпляр логгера
func NewLogger() *Logger {
	return &Logger{
		InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// LogInfo логирует информационное сообщение
func (l *Logger) LogInfo(message string) {
	l.InfoLogger.Println(message) // Вывод в stdout с префиксом INFO
}

// LogError логирует сообщение об ошибке
func (l *Logger) LogError(message string) {
	l.ErrorLogger.Println(message) // Вывод в stderr с префиксом ERROR
}
