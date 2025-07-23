package https

import (
	"encoding/json"
	"net/http"
	"person_service/pkg/logging"
	"person_service/pkg/storage"
	"strconv"
	"strings"
	"unicode"
)

// ручка сохранения в бд
func SaveHandler(storage *storage.Storage, logger *logging.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo("SaveHandler started")

		if r.Method != http.MethodPost {
			logger.LogError("Ну и что мы стучим с неправильным методом? Я жду POST: " + r.Method)
			http.Error(w, "Ну и что мы стучим с неправильным методом? Я жду POST", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Query().Get("ID")
		name := r.URL.Query().Get("Name")
		logger.LogInfo("Received params - ID: " + idStr + ", Name: " + name)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.LogError("Я жду циферку: " + idStr)
			http.Error(w, "Я жду циферку", http.StatusBadRequest)
			return
		}

		// проверка на пустоту
		if len(name) == 0 {
			logger.LogError("Ну пусто же")
			http.Error(w, "Ну пусто же", http.StatusBadRequest)
			return
		}
		// убираю пробелы
		name = strings.ReplaceAll(name, " ", "")

		// Валидация и форматирование имени
		for _, symbol := range name {
			if !(unicode.Is(unicode.Cyrillic, symbol) || unicode.Is(unicode.Latin, symbol)) {
				logger.LogError("Разрешены только буквы кириллицы или латиницы")
				http.Error(w, "Разрешены только буквы кириллицы или латиницы", http.StatusBadRequest)
				return
			}
		}
		formName := formatName(name)

		// Проверка на существующий ID
		if _, exists := storage.Get(id); exists {
			logger.LogError("Такой айдишник уже есть: " + idStr)
			http.Error(w, "Такой айдишник уже есть", http.StatusBadRequest)
			return
		}
		// Проверка на существующее имя
		for _, existingName := range storage.GetAllNames() {
			if existingName == formName {
				logger.LogError("Такое имя уже есть: " + formName)
				http.Error(w, "Такое имя уже есть", http.StatusBadRequest)
				return
			}
		}

		storage.Save(id, formName)
		logger.LogInfo("Successfully saved data")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "Тебя записали"})
	}
}

func formatName(name string) string {
	//	name = strings.ReplaceAll(name, " ", "")
	reName := []rune(name)
	reName[0] = unicode.ToUpper(reName[0])
	for i := 1; i < len(reName); i++ {
		reName[i] = unicode.ToLower(reName[i])

	}

	return string(reName)

}

// ручка чтения из бд
func GetPersonHandler(storage *storage.Storage, logger *logging.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo("GetPersonHandler started")

		if r.Method != http.MethodGet {
			logger.LogError("Ну и что мы стучим с неправильным методом? Я жду GET: " + r.Method)
			http.Error(w, "Ну и что мы стучим с неправильным методом? Я жду GET", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Query().Get("ID")
		logger.LogInfo("Received ID: " + idStr)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.LogError("В параметре я жду циферку: " + idStr)
			http.Error(w, "В параметре я жду циферку", http.StatusBadRequest)
			return
		}

		name, ok := storage.Get(id)
		if !ok {
			logger.LogError("Такого имени нет: " + idStr)
			http.Error(w, "Такого имени нет", http.StatusBadRequest)
			return
		}

		logger.LogInfo("А вот и имя, держи:")
		json.NewEncoder(w).Encode(map[string]string{"name": name})
	}
}
