package datastore

import (
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/gofr"
)

type User interface {
	GetById(ctx *gofr.Context, userId string) (*models.User, error)
	DeleteById(ctx *gofr.Context, userId string) (string, error)
}
