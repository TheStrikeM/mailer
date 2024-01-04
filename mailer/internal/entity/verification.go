package entity

type VerificationBody struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
