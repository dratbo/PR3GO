package api

import (
	"encoding/json"  //Кодирование/декодирование JSON
	"errors"  //Создание и обработка ошибок
	"net/http"  //HTTP сервер и клиент
	"strconv"  //Преобразование строк ↔ числа
	"strings"  //Работа со строками

	"example.com/pz3-http/internal/storage"  //База данных, файлы, кэш
)

type Handlers struct {
	Store *storage.MemoryStore  // Ссылка на хранилище
}

func NewHandlers(store *storage.MemoryStore) *Handlers {
	return &Handlers{Store: store}
}

// GET /tasks
func (h *Handlers) ListTasks(w http.ResponseWriter, r *http.Request) {  // ListTasks - GET /tasks - возвращает список всех задач
	tasks := h.Store.List()  // Получаем все задачи из хранилища

	// Поддержка простых фильтров через query: ?q=text
	q := strings.TrimSpace(r.URL.Query().Get("q"))  // Фильтрация по query параметру ?q=текст
	if q != "" {
		filtered := tasks[:0]  // Создаем новый slice с нулевой емкостью
		for _, t := range tasks {
			if strings.Contains(strings.ToLower(t.Title), strings.ToLower(q)) {  // Поиск подстроки без учета регистра
				filtered = append(filtered, t)
			}
		}
		tasks = filtered
	}

	JSON(w, http.StatusOK, tasks)  // Отправляем JSON ответ	
}

type createTaskRequest struct {
	Title string `json:"title"`
}

// POST /tasks
func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {  // CreateTask - POST /tasks - создает новую задачу
	if r.Header.Get("Content-Type") != "" && !strings.Contains(r.Header.Get("Content-Type"), "application/json") {  // Проверяем Content-Type
		BadRequest(w, "Content-Type must be application/json")
		return
	}

	var req createTaskRequest  // Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, "invalid json: "+err.Error())
		return
	}
	req.Title = strings.TrimSpace(req.Title)   // Валидация
	if req.Title == "" {
		BadRequest(w, "title is required")
		return
	}

	t := h.Store.Create(req.Title)  // Создаем задачу в хранилище
	JSON(w, http.StatusCreated, t)  // 201 Created
}

// GET /tasks/{id} (простой path-парсер без стороннего роутера)
func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	// Ожидаем путь вида /tasks/123
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 {
		NotFound(w, "invalid path")
		return
	}
	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		BadRequest(w, "invalid id")
		return
	}

	t, err := h.Store.Get(id)
	if err != nil {
		if errors.Is(err, errors.New("not found")) {
			NotFound(w, "task not found")
			return
		}
		Internal(w, "unexpected error")
		return
	}
	JSON(w, http.StatusOK, t)
}
