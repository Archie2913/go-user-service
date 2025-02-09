package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/Archie2913/go-user-service/internal/models"
	"github.com/Archie2913/go-user-service/internal/repository"
	"log"
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
	log.Printf("Проверка существования пользователя: %s", email)
	
	// Проверяем, существует ли пользователь
	_, err := s.repo.GetUserByEmail(email)
	if err == nil {
		log.Printf("Пользователь уже существует: %s", email)
		return fmt.Errorf("user with email %s already exists", email)
	}

	log.Printf("Хеширование пароля для пользователя: %s", email)
	// Хешируем пароль
	hashedPassword := sha256.Sum256([]byte(password))

	user := &models.User{
		Email:    email,
		Password: fmt.Sprintf("%x", hashedPassword),
	}

	log.Printf("Создание пользователя в БД: %s", email)
	if err := s.repo.CreateUser(user); err != nil {
		log.Printf("Ошибка создания пользователя в БД: %v", err)
		return err
	}

	log.Printf("Пользователь успешно создан: %s", email)
	return nil
} 