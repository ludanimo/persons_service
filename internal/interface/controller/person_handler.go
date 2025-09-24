package controller

import (
	"encoding/json"
	"net/http"
	"persons_service/internal/entity"
	"persons_service/internal/interface/repository"
	"persons_service/internal/usecase/person"
	"strconv"
)

type PersonHandler struct {
	usecase person.PersonUsecase
}

func NewPersonHandler(usecase person.PersonUsecase) *PersonHandler {
	return &PersonHandler{usecase: usecase}
}

func (h *PersonHandler) SaveHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//h.slog.Info(`SaveHandler started`)

		if r.Method != http.MethodPost {
			//h.slog.Error("Method not allowed: " + r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Query().Get("ID")
		name := r.URL.Query().Get("Name")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			//h.slog.Error("Invalid ID: " + idStr)
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		err = h.usecase.Save(&entity.Person{ID: id, Name: name})
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
		//h.slog.Info("GetHandler started")

		if r.Method != http.MethodGet {
			//h.slog.Error("Method not allowed: " + r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Query().Get("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			//h.slog.Error("Invalid ID: " + idStr)
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		person, err := h.usecase.Get(id)
		if err != nil {
			h.handleError(w, err)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"name": person.Name})
	}
}

func (h *PersonHandler) handleError(w http.ResponseWriter, err error) {
	switch err {
	case person.ErrInvalidID:
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	case person.ErrEmptyName:
		http.Error(w, "Empty name", http.StatusBadRequest)
	case person.ErrInvalidName:
		http.Error(w, "Invalid name characters", http.StatusBadRequest)
	case person.ErrDuplicateID:
		http.Error(w, "Duplicate ID", http.StatusBadRequest)
	case person.ErrDuplicateName:
		http.Error(w, "Duplicate name", http.StatusBadRequest)
	case repository.ErrNotFound:
		http.Error(w, "Not found", http.StatusBadRequest)
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
