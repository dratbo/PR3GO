package storage

import (
	"errors" //работа с ошибками
	"sync"   //параллельное выполнение
)

type Task struct { //Структура задачи
	ID    int64  `json:"id"`    // Уникальный идентификатор
	Title string `json:"title"` // Текст задачи
	Done  bool   `json:"done"`  // Статус выполнения
}

type MemoryStore struct { // MemoryStore - хранилище в оперативной памяти
	mu    sync.RWMutex    // RWMutex для безопасной работы из нескольких горутин
	auto  int64           // Счетчик для автоинкремента ID
	tasks map[int64]*Task // Map для хранения задач [ID] → Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: make(map[int64]*Task), // Инициализация пустой map
	}
}

func (s *MemoryStore) Create(title string) *Task { // Create - создает новую задачу
	s.mu.Lock()         // Блокировка для записи
	defer s.mu.Unlock() // Разблокировка при выходе из функции
	s.auto++            // Увеличиваем счетчик ID
	t := &Task{ID: s.auto, Title: title, Done: false}
	s.tasks[t.ID] = t // Сохраняем в map
	return t
}

func (s *MemoryStore) Get(id int64) (*Task, error) { // Get - получает задачу по ID
	s.mu.RLock() // Блокировка для чтения (много читателей может быть)
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return t, nil
}

func (s *MemoryStore) List() []*Task { // List - возвращает все задачи
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t) // Копируем задачи в slice
	}
	return out
}

// Update - обновляет задачу (частичное обновление)
func (s *MemoryStore) Update(id int64, updates map[string]interface{}) (*Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.tasks[id]
	if !ok {
		return nil, errors.New("not found")
	}

	// Применяем обновления
	if done, exists := updates["done"]; exists {
		if doneBool, ok := done.(bool); ok {
			t.Done = doneBool
		}
	}
	if title, exists := updates["title"]; exists {
		if titleStr, ok := title.(string); ok {
			t.Title = titleStr
		}
	}

	return t, nil
}

// Delete - удаляет задачу по ID
func (s *MemoryStore) Delete(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tasks[id]; !ok {
		return errors.New("not found")
	}

	delete(s.tasks, id)
	return nil
}
