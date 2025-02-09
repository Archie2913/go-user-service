package app

import (
	"net/http"
	"github.com/Archie2913/go-user-service/internal/config"
	"github.com/Archie2913/go-user-service/internal/handler"
	"github.com/Archie2913/go-user-service/internal/repository"
	"github.com/Archie2913/go-user-service/internal/service"
)

// App содержит все зависимости приложения
type App struct {
	cfg     *config.Config
	handler *handler.Handler
	server  *http.Server
}

// NewApp создает новый экземпляр приложения
func NewApp(cfg *config.Config) (*App, error) {
	// Инициализация репозитория
	repo, err := repository.NewRepository(cfg.DB)
	if err != nil {
		return nil, err
	}

	// Инициализация сервиса
	svc := service.NewService(repo)

	// Инициализация обработчиков
	h := handler.NewHandler(svc)

	// Создание сервера
	server := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: h.SetupRoutes(),
	}

	return &App{
		cfg:     cfg,
		handler: h,
		server:  server,
	}, nil
}

// Run запускает приложение
func (a *App) Run() error {
	return a.server.ListenAndServe()
} 