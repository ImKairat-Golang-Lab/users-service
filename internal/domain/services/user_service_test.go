package services

import (
	"context"
	"testing"
	"time"

	mocks "github.com/ImKairat-Golang-Lab/users-service/internal/mocks"
	gomock "github.com/golang/mock/gomock"
)

func TestRegister(t *testing.T) {
	// Создаем мок контроллер
	ctrl := gomock.NewController(t)

	// Создаем мок интерфейсы
	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockClock := mocks.NewMockClock(ctrl)

	// Задаем константное время и что 2 раза должны его получать
	fixedTime := time.Date(2025, 10, 11, 12, 00, 00, 00, time.UTC)
	mockClock.EXPECT().Now().Return(fixedTime).Times(2)

	// Создаем тестовый объект User'а
	ctx := context.Background()
	expectedUser := User{
		Id:           "",
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
		Login:        "test_user",
		CreatedAt:    fixedTime,
		UpdatedAt:    fixedTime,
	}

	// Задаем ожидаемый вызов метода и его результат
	mockRepo.EXPECT().Save(ctx, expectedUser).Return(nil)

	//
	service := NewUserService(mockRepo, mockClock)
	err := service.Register(ctx, "test@example.com", "hashed_password", "test_user")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
