package datastore

import (
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/gofr"
)

type Password interface {
	Create(ctx *gofr.Context, password *models.Password, userId string) (string, error)
	Delete(ctx *gofr.Context, passwordId string) (string, error)
	Update(ctx *gofr.Context, password *models.Password, passwordId string) (string, error)
	Get(ctx *gofr.Context, passwordId string) (*models.Password, error)
	GetByUserId(ctx *gofr.Context, userId string) ([]*models.Password, error)
}
