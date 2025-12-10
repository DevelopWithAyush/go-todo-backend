package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// SwaggerConfig holds swagger configuration
type SwaggerConfig struct {
	BasePath string
	FilePath string
	Title    string
}

// NewSwaggerHandler creates a Fiber v3 compatible Swagger UI handler
func NewSwaggerHandler(cfg *SwaggerConfig) fiber.Handler {
	// Configure http-swagger handler
	handler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("swagger-ui"),
		httpSwagger.PersistAuthorization(true),
	)

	// Use Fiber's adaptor to convert net/http handler to Fiber handler
	return adaptor.HTTPHandler(handler)
}

// RegisterSwagger sets up swagger documentation routes
func RegisterSwagger(app *fiber.App) {
	// Swagger handler for UI and doc.json
	swaggerHandler := NewSwaggerHandler(&SwaggerConfig{
		BasePath: "/swagger",
		Title:    "TODO App API",
	})

	// Register swagger routes - handle all swagger paths
	app.Get("/swagger/*", swaggerHandler)
}
