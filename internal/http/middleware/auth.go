package middleware

import (
	"fmt"
	"time"

	"github.com/developwithayush/go-todo-app/internal/config"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(cfg *config.Config) fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenStr := c.Cookies(cfg.CookieName)
		if tokenStr == "" {
			return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "invalid claims"})
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return c.Status(401).JSON(fiber.Map{"error": "token expired"})
			}
		}

		c.Locals("userID", claims["sub"])

		return c.Next()
	}
}
