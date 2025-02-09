package config

import (
    "encoding/json"
    "os"
)

// Config содержит все конфигурационные параметры приложения
type Config struct {
    Server   ServerConfig   `json:"server"`
    DB       DatabaseConfig `json:"database"`
    Logger   LoggerConfig   `json:"logger"`
}

type ServerConfig struct {
    Port string `json:"port"`
    Host string `json:"host"`
}

type DatabaseConfig struct {
    URL      string `json:"url"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoggerConfig struct {
    Level  string `json:"level"`
    Format string `json:"format"`
}

// Load загружает конфигурацию из JSON файла
func Load() (*Config, error) {
    // Открываем файл конфигурации
    file, err := os.Open("config/config.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Создаем новый конфиг и декодируем JSON
    config := &Config{}
    if err := json.NewDecoder(file).Decode(config); err != nil {
        return nil, err
    }

    return config, nil
}