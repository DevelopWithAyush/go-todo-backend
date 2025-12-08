package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/developwithayush/go-todo-app/internal/config"
	"github.com/developwithayush/go-todo-app/internal/logger"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	authService *Service
	config      *config.Config
	logr        logger.Logger
}

func NewHandler(authService *Service, config *config.Config, logr logger.Logger) *Handler {
	return &Handler{
		authService: authService,
		config:      config,
		logr:        logr,
	}
}

func (h *Handler) SendOTP(c fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
	}

	if err := c.Bind().Body(&body); err != nil || body.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	if err := h.authService.SendOTP(ctx, body.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to send OTP",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "OTP sent successfully",
	})
}

func (h *Handler) VerifyOTP(c fiber.Ctx) error {  
	fmt.Println("VerifyOTP")
	var body struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	if err := c.Bind().Body(&body); err != nil || body.Email == "" || body.OTP == "" {
		h.logr.Error("Invalid request body", logger.Field("error", err), logger.Field("email", body.Email), logger.Field("otp", body.OTP))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()
	
	token, err := h.authService.VerifyOTP(ctx, body.Email, body.OTP)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Failed to verify OTP",
		})
	}

	secure := h.config.Env == "production" || h.config.Env == "prod"

	cookie := fiber.Cookie{
		Name:     h.config.CookieName,
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   secure,
		HTTPOnly: true,
		SameSite: "lax",
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "OTP verified successfully",
		"token":   token,
	})
}
