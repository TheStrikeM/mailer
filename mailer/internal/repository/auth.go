package repository

import (
	"fmt"
	"log/slog"
	"mailer/internal/config"
	"mailer/internal/entity"
	"mailer/pkg/e"
	"net/smtp"
)

type AuthRepository struct {
	log    *slog.Logger
	config *config.Config
}

func NewAuthRepository(log *slog.Logger, cfg *config.Config) *AuthRepository {
	return &AuthRepository{
		log:    log,
		config: cfg,
	}
}

func (ar *AuthRepository) SendEmailVerification(body entity.VerificationBody) error {
	const op = "AuthRepository.SendEmailVerification()"

	email := ar.config.Email
	auth := smtp.PlainAuth("", email.From, email.Password, email.Host)
	fmt.Println("Hello world email")

	err := smtp.SendMail(email.Host+":"+email.Port, auth, email.From, []string{body.Email}, []byte(body.Code))
	if err != nil {
		return e.WrapIfErr(op, err)
	}

	return nil
}
