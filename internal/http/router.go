package http

import (
	"github.com/gofiber/fiber/v3"

	"github.com/developwithayush/go-todo-app/internal/config"
	"github.com/developwithayush/go-todo-app/internal/domain/auth"
	"github.com/developwithayush/go-todo-app/internal/domain/todo"
	"github.com/developwithayush/go-todo-app/internal/domain/user"
	"github.com/developwithayush/go-todo-app/internal/http/middleware"
	"github.com/developwithayush/go-todo-app/internal/logger"
	"github.com/developwithayush/go-todo-app/internal/util"
)

func RegisterRoutes(app *fiber.App, cfg *config.Config, log logger.Logger) {
	// global middleware
	app.Use(middleware.Recover(log))
	app.Use(middleware.Logging(log))
	app.Use(middleware.CORS())

	// health check
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Swagger documentation
	RegisterSwagger(app)

	// deps
	userRepo := user.NewRepository()
	_ = user.NewService(userRepo) // reserved for future extra logic

	mailer, _ := util.NewMailer(cfg)

	authSvc := auth.NewService(cfg, userRepo, mailer)
	authHandler := auth.NewHandler(authSvc, cfg, log)

	todoRepo := todo.NewRepository()
	todoHandler := todo.NewHandler(todoRepo, log)

	api := app.Group("/api/v1")

	// Auth routes
	api.Post("/auth/send-otp", authHandler.SendOTP)
	api.Post("/auth/verify-otp", authHandler.VerifyOTP)

	// Todo routes (protected)
	authMW := middleware.AuthRequired(cfg)
	todoGroup := api.Group("/todos", authMW)
	todoGroup.Get("/", todoHandler.ListTodos)
	todoGroup.Post("/", todoHandler.CreateTodo)
	todoGroup.Put("/:id", todoHandler.UpdateTodo)
	todoGroup.Delete("/:id", todoHandler.DeleteTodo)
}
