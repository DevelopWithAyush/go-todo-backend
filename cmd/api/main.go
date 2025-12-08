package main

import (
	"github.com/developwithayush/go-todo-app/internal/cache"
	"github.com/developwithayush/go-todo-app/internal/config"
	"github.com/developwithayush/go-todo-app/internal/db"
	"github.com/developwithayush/go-todo-app/internal/http"
	"github.com/developwithayush/go-todo-app/internal/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	logr := logger.NewLogger(cfg)

	defer logr.Sync()

	if err := db.InitMongo(cfg, logr); err != nil {
		logr.Fatal("Failed to initialize MongoDB", logger.Field("error", err))
	}

	if err := cache.InitRedis(cfg, logr); err != nil {
		logr.Error("failed to connect redis", logger.Field(
			"error", err))
	}

	app := fiber.New(fiber.Config{
		AppName:   "Go Todo App",
		BodyLimit: 1024 * 1024 * 10, // 10MB
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Internal Server Error",
			})
		},
	})
	http.RegisterRoutes(app, cfg, logr)

	logr.Info("Server is running on port " + cfg.Port)

	if err := app.Listen(":" + cfg.Port); err != nil {
		logr.Fatal("Failed to start server", logger.Field("error", err))
	}
}
