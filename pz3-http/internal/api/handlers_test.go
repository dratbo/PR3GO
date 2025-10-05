package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/pz3-http/internal/storage"
)

func TestHandlers(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)

	t.Run("GET /tasks - empty list", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/tasks", nil)
		w := httptest.NewRecorder()

		h.ListTasks(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}

		var tasks []*storage.Task
		if err := json.Unmarshal(w.Body.Bytes(), &tasks); err != nil {
			t.Fatalf("failed to parse response: %v", err)
		}

		if len(tasks) != 0 {
			t.Errorf("expected empty list, got %d tasks", len(tasks))
		}
	})

	t.Run("POST /tasks - create task", func(t *testing.T) {
		taskData := map[string]string{"title": "Test task"}
		body, _ := json.Marshal(taskData)

		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.CreateTask(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
		}

		var task storage.Task
		if err := json.Unmarshal(w.Body.Bytes(), &task); err != nil {
			t.Fatalf("failed to parse response: %v", err)
		}

		if task.Title != "Test task" {
			t.Errorf("expected title 'Test task', got '%s'", task.Title)
		}
		if task.Done != false {
			t.Errorf("expected done false, got %v", task.Done)
		}
	})

	t.Run("POST /tasks - invalid json", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.CreateTask(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("POST /tasks - title too long", func(t *testing.T) {
		longTitle := string(make([]byte, 141)) // 141 символов
		taskData := map[string]string{"title": longTitle}
		body, _ := json.Marshal(taskData)

		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.CreateTask(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("GET /tasks/{id} - not found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/tasks/999", nil)
		w := httptest.NewRecorder()

		h.GetTask(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("PATCH /tasks/{id} - mark done", func(t *testing.T) {
		// Сначала создаем задачу
		task := store.Create("Test task for update")

		updateData := map[string]bool{"done": true}
		body, _ := json.Marshal(updateData)

		req := httptest.NewRequest("PATCH", "/tasks/"+string(rune(task.ID)), bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.UpdateTask(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}

		var updatedTask storage.Task
		if err := json.Unmarshal(w.Body.Bytes(), &updatedTask); err != nil {
			t.Fatalf("failed to parse response: %v", err)
		}

		if !updatedTask.Done {
			t.Errorf("expected done true, got %v", updatedTask.Done)
		}
	})

	t.Run("DELETE /tasks/{id}", func(t *testing.T) {
		// Сначала создаем задачу
		task := store.Create("Test task for deletion")

		req := httptest.NewRequest("DELETE", "/tasks/"+string(rune(task.ID)), nil)
		w := httptest.NewRecorder()

		h.DeleteTask(w, req)

		if w.Code != http.StatusNoContent {
			t.Errorf("expected status %d, got %d", http.StatusNoContent, w.Code)
		}

		// Проверяем, что задача действительно удалена
		_, err := store.Get(task.ID)
		if err == nil {
			t.Error("expected task to be deleted, but it still exists")
		}
	})
}
