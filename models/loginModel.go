package models

type Login struct {
	Password string `json:"password" validate:"required, min=8, max=100"`
	Email    string `json:"email" validate:"required, min=10, max=10"`
}
