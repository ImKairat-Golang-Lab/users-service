package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ImKairat-Golang-Lab/users-service/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestUserRegister(t *testing.T) {
	// t.Parallel()
	// t.Run()
	ctrl := gomock.NewController(t)

	mockUserService := mocks.NewMockUserService(ctrl)
	mockLogger := mocks.NewMockLogger(ctrl)

	mockUserService.EXPECT().Register(gomock.Any(), "test@example.com", "hashed_password", "test_user").Return(nil)
	mockLogger.EXPECT().Info(gomock.Any(), gomock.Any())

	userHandler := NewUserHandler(mockUserService, mockLogger)
	reqBody := `{"email":"test@example.com","password":"hashed_password","login":"test_user"}`
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	userHandler.UserRegister(w, req)

	res := w.Result()
	// defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected status code 200, got %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error reading response body: %s", err)
	}

	if string(body) != "\"User registered successfully\"\n" {
		t.Fatalf("unexpected body: %s", body)
	}
}
