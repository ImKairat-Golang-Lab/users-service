package services

import (
	"context"
	"testing"
	"time"

	"github.com/ImKairat-Golang-Lab/users-service/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockClock := mocks.NewMockClock(ctrl)

	fixedTime := time.Date(2025, 10, 11, 12, 00, 00, 00, time.UTC)

	mockClock.EXPECT().Now().Return(fixedTime).Times(2)

	ctx := context.Background()
	expectedUser := User{
		Id:           "",
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
		Login:        "test_user",
		CreatedAt:    fixedTime,
		UpdatedAt:    fixedTime,
	}

	mockRepo.EXPECT().Save(ctx, expectedUser).Return(nil)

	service := NewUserService(mockRepo, mockClock)
	err := service.Register(ctx, "test@example.com", "hashed_password", "test_user")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
