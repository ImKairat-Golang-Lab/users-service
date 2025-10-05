package http

import (
	"encoding/json"
	"net/http"

	"github.com/ImKairat-Golang-Lab/users-service/internal/adapters/http/dto"
	"github.com/ImKairat-Golang-Lab/users-service/internal/domain/services"
)


type UserHandler struct {
	service *services.UserService
}

func NewUsUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	// Парсим тела запроса в DTO (Data Transfer Object)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Валидируем тело запроса
	if req.Login == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "All fields are requaired", http.StatusBadRequest)
		return
	}
	// Вызываем доменную логику через порт
	if err := h.service.Register(r.Context(), req.Email, req.Password, req.Login); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Ответ
	w.WriteHeader(http.StatusCreated)
}
