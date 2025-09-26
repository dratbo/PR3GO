package api

import (
	"log"
	"net/http"
	"time"
)

type statusRecorder struct {  // statusRecorder - обертка для ResponseWriter для захвата статуса
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code  // Запоминаем статус
	sr.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {  // Logging - middleware для логирования запросов
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rec, r)  // Передаем управление следующему обработчику
		log.Printf("%s %s %d %v", r.Method, r.URL.Path, rec.status, time.Since(start))  // Логируем после обработки запроса
	})
}
