package http

import (
	"encoding/json"
	"net/http"

	dto "github.com/ImKairat-Golang-Lab/users-service/internal/adapters/http/dto"
	services "github.com/ImKairat-Golang-Lab/users-service/internal/domain/services"
	ports "github.com/ImKairat-Golang-Lab/users-service/internal/ports"
)

type UserHandler struct {
	service *services.UserService
	logger  ports.Logger
}

func NewUserHandler(service *services.UserService, logger ports.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (uh *UserHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	component := "httpHandler/userRegister"

	// Парсим тела запроса в DTO (Data Transfer Object)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		msg := "Invalid request body"
		status_code := http.StatusBadRequest
		fields := map[string]any{
			"component":   component,
			"status_code": status_code,
			"method":      r.Method,
			"endpoint":    r.URL.Path,
			"user_agent":  r.UserAgent(),
		}

		uh.logger.Warn(msg, fields)
		http.Error(w, msg, status_code)
		return
	}
	// Валидируем тело запроса
	if req.Login == "" || req.Email == "" || req.Password == "" {
		msg := "All fields are requaired"
		status_code := http.StatusBadRequest
		fields := map[string]any{
			"component":   component,
			"status_code": status_code,
			"method":      r.Method,
			"endpoint":    r.URL.Path,
			"user_agent":  r.UserAgent(),
		}

		uh.logger.Warn(msg, fields)
		http.Error(w, msg, status_code)
		return
	}
	// Вызываем доменную логику через порт
	if err := uh.service.Register(r.Context(), req.Email, req.Password, req.Login); err != nil {
		msg := err.Error()
		status_code := http.StatusInternalServerError
		fields := map[string]any{
			"component":   component,
			"status_code": status_code,
			"method":      r.Method,
			"endpoint":    r.URL.Path,
			"user_agent":  r.UserAgent(),
		}

		uh.logger.Warn(msg, fields)
		http.Error(w, msg, status_code)
		return
	}
	// Ответ
	w.WriteHeader(http.StatusCreated)
}
