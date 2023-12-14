package models

type Password struct {
	Password_Id    int    `json:"password_id"`
	Password_Name  string `json:"password_name" validate:"required"`
	Password_Value string `json:"password_value" validate:"required"`
	Password_Type  string `json:"password_type"`
	User_Id        int    `json:"user_id"`
}
