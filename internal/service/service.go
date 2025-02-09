package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/Archie2913/go-user-service/internal/models"
	"github.com/Archie2913/go-user-service/internal/repository"
)

// Service содержит бизнес-логику приложения
type Service struct {
	repo *repository.Repository
}

// NewService создает новый сервис
func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Здесь определяются методы бизнес-логики 

func (s *Service) RegisterUser(email, password string) error {
	// Проверяем, существует ли пользователь
	_, err := s.repo.GetUserByEmail(email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", email)
	}

	// Хешируем пароль
	hashedPassword := sha256.Sum256([]byte(password))

	user := &models.User{
		Email:    email,
		Password: fmt.Sprintf("%x", hashedPassword),
	}

	return s.repo.CreateUser(user)
} 