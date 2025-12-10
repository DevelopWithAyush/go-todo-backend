package dto

// SuccessResponse represents a successful API response
// @Description Standard success response wrapper
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data"`
}

// ErrorResponse represents an error API response
// @Description Standard error response wrapper
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error message"`
}

// MessageResponse represents a simple message response
// @Description Simple message response
type MessageResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Operation completed successfully"`
}
