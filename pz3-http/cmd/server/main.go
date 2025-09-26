package main

import (
	"log" 
	"net/http"

	"example.com/pz3-http/internal/api"
	"example.com/pz3-http/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()  // 1. Инициализация хранилища в памяти
	h := api.NewHandlers(store)  // 2. Создание обработчиков API, передаем хранилище

	mux := http.NewServeMux()  // 3. Создание роутера (маршрутизатора)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {  // 4. Регистрация маршрутов:
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Коллекция
	mux.HandleFunc("GET /tasks", h.ListTasks)  // Получить все задачи
	mux.HandleFunc("POST /tasks", h.CreateTask)  // Создать задачу
	// Элемент
	mux.HandleFunc("GET /tasks/", h.GetTask)  // Получить задачу по ID

	// Подключаем логирование
	handler := api.Logging(mux)  // 5. Подключение middleware (цепочка обработки)

	addr := ":8080"  // 6. Запуск сервера
	log.Println("listening on", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
//Сервер запускается на порту 8080
//Запрос проходит через цепочку: CORS → Logging → Роутер → Обработчик
//Каждый middleware добавляет свою функциональность
