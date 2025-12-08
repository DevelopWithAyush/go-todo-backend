package util

import (
	"strconv"

	"github.com/developwithayush/go-todo-app/internal/config"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	dialer *gomail.Dialer
	from   string
}

func NewMailer(cfg *config.Config) (*Mailer, error) {
	port, err := strconv.Atoi(cfg.SMTPPort)
	if err != nil {
		return nil, err
	}

	dialer := gomail.NewDialer(cfg.SMTPHost, port, cfg.SMTPUser, cfg.SMTPPass)
	return &Mailer{
		dialer: dialer,
		from:   cfg.SMTPUser,
	}, nil
}

func (m *Mailer) SendOTP(to, otp string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "OTP for Go Todo App")
	msg.SetBody("text/plain", "Your OTP is "+otp)
	return m.dialer.DialAndSend(msg)
}
