package usecase

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"mailer/internal/entity"
	"mailer/pkg/e"
)

type AuthUsecase struct {
	log            *slog.Logger
	authRepository AuthRepository
}

func NewAuthUsecase(log *slog.Logger, authRepository AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		log:            log,
		authRepository: authRepository,
	}
}

type AuthRepository interface {
	SendEmailVerification(body entity.VerificationBody) error
}

func (au *AuthUsecase) SendVerification(msg []byte) (err error) {
	const op = "AuthUsecase.SendVerification()"
	defer func() { err = e.WrapIfErr(op, err) }()

	var body entity.VerificationBody
	err = json.Unmarshal(msg, &body)
	fmt.Println(body)
	if err != nil {
		return err
	}

	err = au.authRepository.SendEmailVerification(body)
	if err != nil {
		return err
	}
	return nil
}
