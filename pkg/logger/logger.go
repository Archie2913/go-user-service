package logger

// Logger предоставляет функционал логирования
type Logger struct {
    // конфигурация логгера
}

// NewLogger создает новый логгер
func NewLogger() *Logger {
    return &Logger{}
}

// Info логирует информационное сообщение
func (l *Logger) Info(msg string) {
    // реализация
}

// Error логирует ошибку
func (l *Logger) Error(msg string, err error) {
    // реализация
} 