package handlers

import (
	"github.com/DevanshBatra20-PasswordManager/datastore"
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type authHandler struct {
	authDatastore datastore.Auth
}

func NewAuth(a datastore.Auth) authHandler {
	return authHandler{authDatastore: a}
}

func (a authHandler) Signup(ctx *gofr.Context) (interface{}, error) {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		ctx.Logger.Errorf("Error in binding user: %v", err)
		return nil, errors.InvalidParam{Param: []string{"user"}}
	}
	resp, err := a.authDatastore.Signup(ctx, &user)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a authHandler) Login(ctx *gofr.Context) (interface{}, error) {
	var userCredentials models.Login

	if err := ctx.Bind(&userCredentials); err != nil {
		ctx.Logger.Errorf("Error in binding user: %v", err)
		return nil, errors.InvalidParam{Param: []string{"user"}}
	}
	resp, err := a.authDatastore.Login(ctx, &userCredentials)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
