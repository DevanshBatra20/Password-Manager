package handlers

import (
	"github.com/DevanshBatra20-PasswordManager/datastore"
	"github.com/DevanshBatra20-PasswordManager/helpers"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type userHandler struct {
	userStore datastore.User
}

func NewUser(u datastore.User) userHandler {
	return userHandler{userStore: u}
}

func (u userHandler) GetById(ctx *gofr.Context, userId string) (interface{}, error) {

	if userId == "" {
		return nil, errors.MissingParam{Param: []string{"userId"}}
	}

	if _, err := helpers.ValidateId(userId); err != nil {
		return nil, errors.InvalidParam{Param: []string{"userId"}}
	}

	resp, err := u.userStore.GetById(ctx, userId)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "User",
			ID:     userId,
		}
	}

	return resp, nil
}

func (u userHandler) DeleteById(ctx *gofr.Context, userId string) (interface{}, error) {

	if userId == "" {
		return nil, errors.MissingParam{Param: []string{"userId"}}
	}

	if _, err := helpers.ValidateId(userId); err != nil {
		return nil, errors.InvalidParam{Param: []string{"userId"}}
	}

	resp, err := u.userStore.DeleteById(ctx, userId)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "User",
			ID:     userId,
		}
	}

	return resp, nil
}
