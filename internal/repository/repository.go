package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/Archie2913/go-user-service/internal/config"
	"github.com/Archie2913/go-user-service/internal/models"
	"time"
)

// Repository работает с хранилищем данных
type Repository struct {
	db *sql.DB
}

// NewRepository создает новый репозиторий
func NewRepository(cfg config.DatabaseConfig) (*Repository, error) {
	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

// Здесь определяются методы для работы с базой данных

func (r *Repository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, password, created_at)
		VALUES ($1, $2, $3)
		RETURNING id`

	return r.db.QueryRow(
		query,
		user.Email,
		user.Password,
		time.Now(),
	).Scan(&user.ID)
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