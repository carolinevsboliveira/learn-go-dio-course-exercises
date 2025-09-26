package todo_list_api

import (
	"net/http"
	"strings"
)

func SetupRoutes(handler *TodoHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthCheck)

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetAllTodos(w, r)
		case http.MethodPost:
			handler.CreateTodo(w, r)
		default:
			handler.MethodNotAllowedHandler(w, r)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/todos/")
		if path == "" {
			http.Redirect(w, r, "/todos", http.StatusMovedPermanently)
			return
		}

		switch r.Method {
		case http.MethodGet:
			handler.GetTodoByID(w, r)
		case http.MethodPut:
			handler.UpdateTodo(w, r)
		case http.MethodDelete:
			handler.DeleteTodo(w, r)
		default:
			handler.MethodNotAllowedHandler(w, r)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{
				"success": true,
				"message": "Todo API v1.0.0",
				"endpoints": {
					"GET /health": "Health check",
					"GET /todos": "Get all todos",
					"POST /todos": "Create a new todo",
					"GET /todos/{id}": "Get todo by ID",
					"PUT /todos/{id}": "Update todo by ID",
					"DELETE /todos/{id}": "Delete todo by ID"
				},
				"example": {
					"create_todo": {
						"method": "POST",
						"url": "/todos",
						"body": {
							"title": "Learn Go",
							"description": "Study Go programming language"
						}
					}
				}
			}`, http.StatusOK)
			return
		}

		handler.NotFoundHandler(w, r)
	})

	return mux
}
