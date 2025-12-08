package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/developwithayush/go-todo-app/internal/config"
	"github.com/developwithayush/go-todo-app/internal/domain/user"
	"github.com/developwithayush/go-todo-app/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	userRepo user.Repository
	mailer *util.Mailer
	config *config.Config
}

 


func NewService(cfg *config.Config, userRepo user.Repository, mailer *util.Mailer) *Service {
	return &Service{
		config: cfg,
		userRepo: userRepo,
		mailer: mailer,
	}
}


func (s *Service) SendOTP(ctx context.Context, email string) error {
	otp := util.GenerateOTP()
	hashedOTP, err := util.HashOTP(otp)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(10 * time.Minute)
	user, err := s.userRepo.UpsertOTP(ctx, email, hashedOTP, expiresAt)
	if err != nil {
		return err
	}
	return s.mailer.SendOTP(user.Email, otp) 
}


func (s *Service) VerifyOTP(ctx context.Context, email, otp string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user not found")
	} 

	if time.Now().After(user.OTPExpiresAt) { 
		fmt.Println("otp expired")
		return " ", errors.New("otp expired")
	}
	
	if !util.CheckOTP(user.OTPHash, otp) {
		fmt.Println("invalid otp")
		return "", errors.New("invalid otp")
	}

	_ = s.userRepo.ClearOTP(ctx, user.ID)
	

	claims := jwt.MapClaims{
		"sub": user.ID.Hex(),
		"exp": time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}