package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {  // ErrorResponse - стандартная структура ошибки
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, v any) {  // JSON - универсальная функция для отправки JSON ответов
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // Важно для русских символов!
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)  // Кодируем в JSON и отправляем
}

func BadRequest(w http.ResponseWriter, msg string) {  // Вспомогательные функции для стандартных ошибок
	JSON(w, http.StatusBadRequest, ErrorResponse{Error: msg})
}

func NotFound(w http.ResponseWriter, msg string) {  // Вспомогательные функции для стандартных ошибок
	JSON(w, http.StatusNotFound, ErrorResponse{Error: msg})
}

func Internal(w http.ResponseWriter, msg string) {  // Вспомогательные функции для стандартных ошибок
	JSON(w, http.StatusInternalServerError, ErrorResponse{Error: msg})
}
