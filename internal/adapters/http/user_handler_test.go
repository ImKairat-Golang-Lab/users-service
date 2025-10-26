package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ImKairat-Golang-Lab/users-service/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestUserRegister(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserService := mocks.NewMockUserService(ctrl)
	mockLogger := mocks.NewMockLogger(ctrl)
	ctx := context.Background()

	mockUserService.EXPECT().Register(ctx, "test@example.com", "hashed_password", "test_user").Return(nil)

	userHandler := NewUserHandler(mockUserService, mockLogger)
	req := httptest.NewRequest(http.MethodPost, "/register", nil)
	w := httptest.NewRecorder()

	userHandler.UserRegister(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code 200, got %d", res.StatusCode)
	}

	body := make([]byte, res.ContentLength)
	if _, err := res.Body.Read(body); err != nil {
		t.Fatalf("error with read response body: %s", err)
	}
	if string(body) != "User registered successfully" {
		t.Fatalf("unexpected body: %s", body)
	}
}
