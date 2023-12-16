package datastore

import (
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/gofr"
)

type Auth interface {
	Signup(ctx *gofr.Context, user *models.User) (string, error)
	Login(ctx *gofr.Context, userCredentials *models.Login) (*models.Signup, error)
}
