package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/developwithayush/go-todo-app/internal/logger"
)

func Recover(log logger.Logger) fiber.Handler {
	return func(c fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic recovered", logger.Field("panic", r))
				_ = c.Status(500).JSON(fiber.Map{"error": "internal server error"})
			}
		}()
		return c.Next()
	}
}
