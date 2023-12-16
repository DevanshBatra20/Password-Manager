package datastore

import (
	"github.com/DevanshBatra20-PasswordManager/exception"
	"github.com/DevanshBatra20-PasswordManager/helpers"
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/gofr"
)

type auth struct{}

func NewAuth() *auth {
	return &auth{}
}

func (a *auth) Signup(ctx *gofr.Context, user *models.User) (string, error) {
	var count int
	err := ctx.DB().QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE email=(?) || phone=(?)",
		user.Email, user.Phone).Scan(&count)
	if err != nil {
		return "Error in signing up", err
	}
	if count > 0 {
		return "Error in signing up", exception.UserAlreadyExists{Email: *user.Email, Phone: *user.Phone}
	}
	token, err := helpers.GenerateJwtToken(*user.Email, *user.First_Name, *user.Last_Name)
	if err != nil {
		return "Error in signing up", err
	}
	hashPassword, err := helpers.HashPassword(*user.Password)
	if err != nil {
		return "Error in signing up", err
	}
	_, err = ctx.DB().ExecContext(ctx, "INSERT INTO users (first_name, last_name, phone, email, token, password) VALUES(?, ?, ?, ?, ?, ?)",
		user.First_Name, user.Last_Name, user.Phone, user.Email, token, hashPassword)
	if err != nil {
		return "Error in signing up", err
	}

	return "Signed up sucessfully", nil
}

func (a *auth) Login(ctx *gofr.Context, userCredentials *models.Login) (*models.Signup, error) {
	var foundUser models.Signup

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM users WHERE email=(?)", userCredentials.Email).
		Scan(&foundUser.ID, &foundUser.First_Name, &foundUser.Last_Name,
			&foundUser.Phone, &foundUser.Email, &foundUser.Token, &foundUser.Password)
	if err != nil {
		return &models.Signup{}, exception.UserNotFound{Email: *userCredentials.Email}
	}

	isPasswordValid := helpers.VerifyPassword(*foundUser.Password, *userCredentials.Password)
	if !isPasswordValid {
		return &models.Signup{}, exception.InvalidCredentials{
			Password: *userCredentials.Password,
			Email:    *userCredentials.Email,
		}
	}

	token, _ := helpers.GenerateJwtToken(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name)
	_, err = ctx.DB().ExecContext(ctx, "UPDATE users SET token=(?) WHERE email=(?)",
		token, foundUser.Email)
	if err != nil {
		return &models.Signup{}, err
	}

	return &foundUser, nil
}
