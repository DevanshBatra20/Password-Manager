package datastore

import (
	"github.com/DevanshBatra20-PasswordManager/helpers"
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/gofr"
)

func Signup(c *gofr.Context, user *models.User) (string, error) {
	var count int
	err := c.DB().QueryRowContext(c, "SELECT COUNT(*) FROM users WHERE email=(?) || phone=(?)",
		user.Email, user.Phone).Scan(&count)
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "User with the same email or phone already exists", nil
	}
	token, err := helpers.GenerateJwtToken(*user.Email, *user.First_Name, *user.Last_Name)
	if err != nil {
		return "", err
	}
	hashPassword, err := helpers.HashPassword(*user.Password)
	if err != nil {
		return "", err
	}
	_, err = c.DB().ExecContext(c, "INSERT INTO users (first_name, last_name, phone, email, token, password) VALUES(?, ?, ?, ?, ?, ?)",
		user.First_Name, user.Last_Name, user.Phone, user.Email, token, hashPassword)
	if err != nil {
		return "", err
	}
	return "Signup completed", nil
}
