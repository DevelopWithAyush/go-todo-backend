package dto

import "time"

// CreateTodoRequest represents the request body for creating a todo
// @Description Request body for creating a new todo item
type CreateTodoRequest struct {
	Title       string `json:"title" example:"Buy groceries" validate:"required"`
	Description string `json:"description" example:"Milk, eggs, bread"`
}

// UpdateTodoRequest represents the request body for updating a todo
// @Description Request body for updating an existing todo item
type UpdateTodoRequest struct {
	Title       string `json:"title" example:"Buy groceries (updated)"`
	Description string `json:"description" example:"Milk, eggs, bread, butter"`
}

// TodoResponse represents a single todo item in the response
// @Description Todo item response structure
type TodoResponse struct {
	ID          string    `json:"id" example:"507f1f77bcf86cd799439011"`
	UserID      string    `json:"userId" example:"507f1f77bcf86cd799439012"`
	Title       string    `json:"title" example:"Buy groceries"`
	Description string    `json:"description" example:"Milk, eggs, bread"`
	Completed   bool      `json:"completed" example:"false"`
	Position    int       `json:"position" example:"0"`
	CreatedAt   time.Time `json:"createdAt" example:"2024-01-15T10:30:00Z"`
	UpdatedAt   time.Time `json:"updatedAt" example:"2024-01-15T10:30:00Z"`
}

// TodoListResponse represents the response containing list of todos
// @Description Response containing list of todo items
type TodoListResponse struct {
	Success bool           `json:"success" example:"true"`
	Data    []TodoResponse `json:"data"`
}

// TodoCreateResponse represents the response after creating a todo
// @Description Response after successfully creating a todo
type TodoCreateResponse struct {
	Success bool         `json:"success" example:"true"`
	Data    TodoResponse `json:"data"`
}
