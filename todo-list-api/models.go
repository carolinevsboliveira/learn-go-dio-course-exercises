package todo_list_api

import (
	"time"
)

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
}

type TodoStore interface {
	GetAll() []Todo
	GetByID(id int) (*Todo, error)
	Create(todo Todo) Todo
	Update(id int, updates UpdateTodoRequest) (*Todo, error)
	Delete(id int) error
}

type InMemoryTodoStore struct {
	todos  []Todo
	nextID int
}

func NewInMemoryTodoStore() *InMemoryTodoStore {
	return &InMemoryTodoStore{
		todos:  make([]Todo, 0),
		nextID: 1,
	}
}

func (s *InMemoryTodoStore) GetAll() []Todo {
	return s.todos
}

func (s *InMemoryTodoStore) GetByID(id int) (*Todo, error) {
	for _, todo := range s.todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, ErrTodoNotFound
}

func (s *InMemoryTodoStore) Create(todo Todo) Todo {
	todo.ID = s.nextID
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	s.todos = append(s.todos, todo)
	s.nextID++
	return todo
}

func (s *InMemoryTodoStore) Update(id int, updates UpdateTodoRequest) (*Todo, error) {
	for i, todo := range s.todos {
		if todo.ID == id {
			if updates.Title != nil {
				s.todos[i].Title = *updates.Title
			}
			if updates.Description != nil {
				s.todos[i].Description = *updates.Description
			}
			if updates.Completed != nil {
				s.todos[i].Completed = *updates.Completed
			}
			s.todos[i].UpdatedAt = time.Now()
			return &s.todos[i], nil
		}
	}
	return nil, ErrTodoNotFound
}

func (s *InMemoryTodoStore) Delete(id int) error {
	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return ErrTodoNotFound
}
