package http

import (
	"encoding/json"
	"net/http"

	dto "github.com/ImKairat-Golang-Lab/users-service/internal/adapters/http/dto"
	ports "github.com/ImKairat-Golang-Lab/users-service/internal/ports"
)

type UserHandler struct {
	service ports.UserService
	logger  ports.Logger
}

func NewUserHandler(service ports.UserService, logger ports.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (uh *UserHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			uh.logger.Warn("failed to close request body", map[string]any{
				"error": err.Error(),
			})
		}
	}()
	component := "httpHandler/userRegister"
	var req dto.RegisterRequest

	//
	if r.Method != http.MethodPost {
		msg := "method not allowed"
		status_code := http.StatusMethodNotAllowed

		uh.logger.Warn(msg, logFields(r, component, status_code))
		writeJSON(w, uh.logger, status_code, dto.ErrorResponse{Error: msg})
		return
	}

	// Парсим тела запроса в DTO (Data Transfer Object)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		status_code := http.StatusBadRequest

		uh.logger.Warn(err.Error(), logFields(r, component, status_code))
		writeJSON(w, uh.logger, status_code, dto.ErrorResponse{Error: err.Error()})
		return
	}
	// Валидируем тело запроса
	if req.Login == "" || req.Email == "" || req.Password == "" {
		msg := "All fields are requaired"
		status_code := http.StatusBadRequest

		uh.logger.Warn(msg, logFields(r, component, status_code))
		writeJSON(w, uh.logger, status_code, dto.ErrorResponse{Error: msg})
		return
	}
	// Вызываем доменную логику через порт
	if err := uh.service.Register(r.Context(), req.Email, req.Password, req.Login); err != nil {
		status_code := http.StatusInternalServerError

		uh.logger.Warn(err.Error(), logFields(r, component, status_code))
		writeJSON(w, uh.logger, status_code, dto.ErrorResponse{Error: err.Error()})
		return
	}
	// Ответ
	msg := "User registered successfully"
	status_code := http.StatusCreated
	uh.logger.Info(msg, logFields(r, component, status_code))
	writeJSON(w, uh.logger, status_code, msg)
}
