package todo

import (
	"context"
	"time"

	"github.com/developwithayush/go-todo-app/internal/db"
	"github.com/developwithayush/go-todo-app/internal/logger"
	"github.com/developwithayush/go-todo-app/internal/util"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	repo Repository
	logr logger.Logger
}

func NewHandler(repo Repository, logr logger.Logger) *Handler {
	return &Handler{
		repo: repo,
		logr: logr,
	}
}

// get all todos for a user
func (h *Handler) ListTodos(c fiber.Ctx) error {
	userIdString, ok := c.Locals("userID").(string)
	if !ok {
		return util.Error(c, fiber.StatusUnauthorized, "Invalid user session")
	}
	userID, err := primitive.ObjectIDFromHex(userIdString)
	if err != nil {
		return util.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	todos, err := h.repo.ListByUser(ctx, userID)
	if err != nil {
		return util.Error(c, fiber.StatusInternalServerError, "Failed to list todos")
	}

	return util.OK(c, todos)
}

// create a new todo
func (h *Handler) CreateTodo(c fiber.Ctx) error {
	userIdString, ok := c.Locals("userID").(string)
	if !ok {
		return util.Error(c, fiber.StatusUnauthorized, "Invalid user session")
	}
	userID, err := primitive.ObjectIDFromHex(userIdString)
	if err != nil {
		return util.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.Bind().Body(&body); err != nil || body.Title == "" {
		return util.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var last Todo
	opts := options.FindOne().SetSort(bson.M{"position": -1})
	err = db.Todos.FindOne(ctx, bson.M{"userId": userID}, opts).Decode(&last)
	position := 0
	if err == nil {
		position = last.Position + 1
	}

	todo := Todo{
		ID:          primitive.NewObjectID(),
		UserID:      userID,
		Title:       body.Title,
		Description: body.Description,
		Completed:   false,
		Position:    position,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = h.repo.Create(ctx, todo)
	if err != nil {
		return util.Error(c, fiber.StatusInternalServerError, "Failed to create todo")
	}

	return util.OK(c, todo)
}

// update a todo
func (h *Handler) UpdateTodo(c fiber.Ctx) error {
	userIdString, ok := c.Locals("userID").(string)
	if !ok {
		return util.Error(c, fiber.StatusUnauthorized, "Invalid user session")
	}
	userID, err := primitive.ObjectIDFromHex(userIdString)
	if err != nil {
		return util.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	todoID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return util.Error(c, fiber.StatusBadRequest, "Invalid todo ID")
	}
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.Bind().Body(&body); err != nil || body.Title == "" {
		return util.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	update := bson.M{"updatedAt": time.Now()}
	if body.Title != "" {
		update["title"] = body.Title
	}
	if body.Description != "" {
		update["description"] = body.Description
	}
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()
	if err := h.repo.Update(ctx, userID, todoID, update); err != nil {
		return util.Error(c, fiber.StatusInternalServerError, "Failed to update todo")
	}

	return util.OK(c, "Todo updated successfully")
}

// delete a todo
func (h *Handler) DeleteTodo(c fiber.Ctx) error {
	userIdString, ok := c.Locals("userID").(string)
	if !ok {
		return util.Error(c, fiber.StatusUnauthorized, "Invalid user session")
	}
	userID, err := primitive.ObjectIDFromHex(userIdString)
	if err != nil {
		return util.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	todoID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return util.Error(c, fiber.StatusBadRequest, "Invalid todo ID")
	}
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()
	if err := h.repo.Delete(ctx, userID, todoID); err != nil {
		return util.Error(c, fiber.StatusInternalServerError, "Failed to delete todo")
	}
	return util.OK(c, "Todo deleted successfully")
} 

