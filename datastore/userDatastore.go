package datastore

import (
	"database/sql"

	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type user struct{}

func NewUser() *user {
	return &user{}
}

func (u *user) GetById(ctx *gofr.Context, userId string) (*models.User, error) {
	var user models.User

	err := ctx.DB().QueryRowContext(ctx, "SELECT id, first_name, last_name, phone, email, password FROM users WHERE id = (?)", userId).
		Scan(&user.ID, &user.First_Name, &user.Last_Name, &user.Phone, &user.Email, &user.Password)
	switch err {
	case sql.ErrNoRows:
		return &models.User{}, errors.EntityNotFound{Entity: "User", ID: userId}
	case nil:
		return &user, nil
	default:
		return &models.User{}, err
	}
}

func (u *user) DeleteById(ctx *gofr.Context, userId string) (string, error) {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM users WHERE id = (?)", userId)
	switch err {
	case sql.ErrNoRows:
		return "Error in deleting user", errors.EntityNotFound{Entity: "User", ID: userId}
	case nil:
		return "User deleted successfully", nil
	default:
		return "Error in deleteing user", err
	}
}
