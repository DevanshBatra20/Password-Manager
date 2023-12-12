package models

type User struct {
	ID         int    `json:"id"`
	First_Name string `json:"first_name" validate:"required, min=2, max=100"`
	Last_Name  string `json:"last_name" validate:"required, min=2, max=100"`
	Phone      string `json:"phone" validate:"required, min=2, max=100"`
	Email      string `json:"email" validate:"required, min=10, max=10"`
	Password   string `json:"password" validate:"required, min=8, max=100"`
}
