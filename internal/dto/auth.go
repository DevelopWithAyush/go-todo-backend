package dto

// SendOTPRequest represents the request body for sending OTP
// @Description Request body for sending OTP to user's email
type SendOTPRequest struct {
	Email string `json:"email" example:"user@example.com" validate:"required,email"`
}

// VerifyOTPRequest represents the request body for verifying OTP
// @Description Request body for verifying OTP
type VerifyOTPRequest struct {
	Email string `json:"email" example:"user@example.com" validate:"required,email"`
	OTP   string `json:"otp" example:"1234" validate:"required,len=4"`
}

// VerifyOTPResponse represents the response after successful OTP verification
// @Description Response containing JWT token after successful OTP verification
type VerifyOTPResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"OTP verified successfully"`
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
