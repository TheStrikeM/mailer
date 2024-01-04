package delivery

type AuthUsecase interface {
	SendVerification([]byte) error
}
