package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"example.com/pz3-http/internal/storage"
)

type Handlers struct {
	Store *storage.MemoryStore
}

func NewHandlers(store *storage.MemoryStore) *Handlers {
	return &Handlers{Store: store}
}

// GET /tasks
func (h *Handlers) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.Store.List()

	// Поддержка простых фильтров через query: ?q=text
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if q != "" {
		filtered := tasks[:0]
		for _, t := range tasks {
			if strings.Contains(strings.ToLower(t.Title), strings.ToLower(q)) {
				filtered = append(filtered, t)
			}
		}
		tasks = filtered
	}

	JSON(w, http.StatusOK, tasks)
}

type createTaskRequest struct {
	Title string `json:"title"`
}

// POST /tasks
func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "" && !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		BadRequest(w, "Content-Type must be application/json")
		return
	}

	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, "invalid json: "+err.Error())
		return
	}
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		BadRequest(w, "title is required")
		return
	}

	// Валидация длины (1-140 символов)
	if len(req.Title) < 1 || len(req.Title) > 140 {
		BadRequest(w, "title must be between 1 and 140 characters")
		return
	}

	t := h.Store.Create(req.Title)
	JSON(w, http.StatusCreated, t)
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

type updateTaskRequest struct {
	Title *string `json:"title,omitempty"` // Указатель для различения отсутствия поля и пустого значения
	Done  *bool   `json:"done,omitempty"`
}

// PATCH /tasks/{id}
func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "" && !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		BadRequest(w, "Content-Type must be application/json")
		return
	}

	// Парсим ID из URL
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

	var req updateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, "invalid json: "+err.Error())
		return
	}

	// Собираем обновления
	updates := make(map[string]interface{})
	if req.Title != nil {
		title := strings.TrimSpace(*req.Title)
		if title == "" {
			BadRequest(w, "title cannot be empty")
			return
		}
		if len(title) < 1 || len(title) > 140 {
			BadRequest(w, "title must be between 1 and 140 characters")
			return
		}
		updates["title"] = title
	}
	if req.Done != nil {
		updates["done"] = *req.Done
	}

	// Если нет полей для обновления
	if len(updates) == 0 {
		BadRequest(w, "no fields to update")
		return
	}

	t, err := h.Store.Update(id, updates)
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

// DELETE /tasks/{id}
func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Парсим ID из URL
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

	err = h.Store.Delete(id)
	if err != nil {
		if errors.Is(err, errors.New("not found")) {
			NotFound(w, "task not found")
			return
		}
		Internal(w, "unexpected error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
