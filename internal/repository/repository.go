package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/Archie2913/go-user-service/internal/config"
	"github.com/Archie2913/go-user-service/internal/models"
	"time"
	"log"
)

// Repository работает с хранилищем данных
type Repository struct {
	db *sql.DB
}

// NewRepository создает новый репозиторий
func NewRepository(cfg config.DatabaseConfig) (*Repository, error) {
	var db *sql.DB
	var err error
	
	// Попытки подключения к БД
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", cfg.URL)
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}

		// Проверяем подключение
		if err = db.Ping(); err != nil {
			time.Sleep(time.Second * 1)
			continue
		}

		// Если подключились успешно, выходим из цикла
		break
	}

	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

// Здесь определяются методы для работы с базой данных

func (r *Repository) CreateUser(user *models.User) error {
	log.Printf("Выполнение запроса на создание пользователя: %s", user.Email)
	
	query := `
		INSERT INTO users (email, password, created_at)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := r.db.QueryRow(
		query,
		user.Email,
		user.Password,
		time.Now(),
	).Scan(&user.ID)

	if err != nil {
		log.Printf("Ошибка выполнения запроса: %v", err)
		return err
	}

	log.Printf("Пользователь успешно создан в БД: %s (ID: %d)", user.Email, user.ID)
	return nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password, created_at
		FROM users
		WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
} 