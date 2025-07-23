package storage

import (
	"person_service/pkg/logging"
	"sync"
)

// Storage хранилище имен/айдишников
type Storage struct {
	mu     sync.RWMutex
	names  map[int]string // Мапа для хранения данных (id = name)
	logger *logging.Logger
}

// Создаем новое хранилище
func NewStorage(logger *logging.Logger) *Storage {
	return &Storage{
		names:  make(map[int]string),
		logger: logger,
	}
}

// Save сохраняет имя по ID
func (s *Storage) Save(id int, name string) {
	s.mu.Lock()         // Блокировка на запись
	defer s.mu.Unlock() // Гарантированная разблокировка

	s.names[id] = name // Сохранение значения
	s.logger.LogInfo(  // Логгирование события
		"Saved name: " + name + " with ID: " + string(id),
	)
}

// Get возвращает имя по ID
func (s *Storage) Get(id int) (string, bool) {
	s.mu.RLock()         // Блокировка на чтение
	defer s.mu.RUnlock() // Гарантированная разблокировка

	name, ok := s.names[id] // Получение значения
	if ok {
		s.logger.LogInfo(
			"Retrieved name: " + name + " for ID: " + string(id),
		)
	} else {
		s.logger.LogError(
			"Name not found for ID: " + string(id),
		)
	}
	return name, ok
}

// возвращение всего списка для проверки дублей
func (s *Storage) GetAllNames() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	names := make([]string, 0, len(s.names))
	for _, name := range s.names {
		names = append(names, name)
	}
	return names
}
