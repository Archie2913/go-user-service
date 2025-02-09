package main

import (
    "log"
    
    "github.com/Archie2913/go-user-service/internal/app"
    "github.com/Archie2913/go-user-service/internal/config"
)

func main() {
    // Загрузка конфигурации
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Инициализация приложения
    app, err := app.NewApp(cfg)
    if err != nil {
        log.Fatalf("Failed to initialize app: %v", err)
    }
    
    // Запуск приложения
    if err := app.Run(); err != nil {
        log.Fatalf("Error running app: %v", err)
    }
} 