package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/developwithayush/go-todo-app/internal/logger"
)

func Logging(log logger.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		log.Info("request",
			logger.Field("method", c.Method()),
			logger.Field("path", c.Path()),
			logger.Field("status", c.Response().StatusCode()),
			logger.Field("duration", duration.String()),
		)

		return err
	}
}
