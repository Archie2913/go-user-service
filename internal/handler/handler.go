package handler

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/Archie2913/go-user-service/internal/service"
)

// Handler обрабатывает HTTP-запросы
type Handler struct {
	service *service.Service
}

// NewHandler создает новый обработчик
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	log.Printf("Получен запрос на регистрацию")
	
	if r.Method != http.MethodPost {
		log.Printf("Неверный метод: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Ошибка декодирования JSON: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Попытка регистрации пользователя: %s", req.Email)

	if err := h.service.RegisterUser(req.Email, req.Password); err != nil {
		log.Printf("Ошибка регистрации пользователя: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Пользователь успешно зарегистрирован: %s", req.Email)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", h.Register)
	return mux
}

// Здесь определяются методы для обработки HTTP-запросов 