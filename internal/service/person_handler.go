package service

import (
	"encoding/json"
	"net/http"
	"persons_service/internal/entity"
	"strconv"
)

type PersonHandler struct {
	service *entity.PersonService
}

func NewPersonHandler(service *entity.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

func (h *PersonHandler) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Query().Get("ID")
		name := r.URL.Query().Get("Name")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		err = h.service.Save(&entity.Person{ID: id, Name: name})
		if err != nil {
			h.handleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "Ok"})
	}
}

func (h *PersonHandler) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Query().Get("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		person, err := h.service.Get(id)
		if err != nil {
			h.handleError(w, err)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"name": person.Name})
	}
}

func (h *PersonHandler) handleError(w http.ResponseWriter, err error) {
	switch err {
	case entity.ErrInvalidID:
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	case entity.ErrEmptyName:
		http.Error(w, "Empty name", http.StatusBadRequest)
	case entity.ErrInvalidName:
		http.Error(w, "Invalid name characters", http.StatusBadRequest)
	case entity.ErrDuplicateID:
		http.Error(w, "Duplicate ID", http.StatusBadRequest)
	case entity.ErrDuplicateName:
		http.Error(w, "Duplicate name", http.StatusBadRequest)
	case entity.ErrNotFound:
		http.Error(w, "Not found", http.StatusBadRequest)
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
