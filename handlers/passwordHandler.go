package handlers

import (
	"github.com/DevanshBatra20-PasswordManager/datastore"
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type passwordHandler struct {
	passwordStore datastore.Password
}

func NewPassword(p datastore.Password) passwordHandler {
	return passwordHandler{passwordStore: p}
}

func (p *passwordHandler) Create(ctx *gofr.Context) (interface{}, error) {
	userId := ctx.PathParam("userId")
	var password models.Password
	if userId == "" {
		return nil, errors.MissingParam{Param: []string{"userId"}}
	}

	if err := ctx.Bind(&password); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	resp, err := p.passwordStore.Create(ctx, &password, userId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *passwordHandler) Delete(ctx *gofr.Context) (interface{}, error) {
	passwordId := ctx.PathParam("passwordId")
	if passwordId == "" {
		return nil, errors.MissingParam{Param: []string{"passwordId"}}
	}

	resp, err := p.passwordStore.Delete(ctx, passwordId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *passwordHandler) Update(ctx *gofr.Context) (interface{}, error) {
	passwordId := ctx.PathParam("passwordId")
	var password models.Password
	if passwordId == "" {
		return nil, errors.MissingParam{Param: []string{"passwordId"}}
	}

	if err := ctx.Bind(&password); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := p.passwordStore.Update(ctx, &password, passwordId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *passwordHandler) Get(ctx *gofr.Context) (interface{}, error) {
	passwordId := ctx.PathParam("passwordId")
	if passwordId == "" {
		return nil, errors.MissingParam{Param: []string{"passwordId"}}
	}

	resp, err := p.passwordStore.Get(ctx, passwordId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *passwordHandler) GetByUserId(ctx *gofr.Context) (interface{}, error) {
	userId := ctx.PathParam("userId")
	if userId == "" {
		return nil, errors.MissingParam{Param: []string{"userId"}}
	}

	resp, err := p.passwordStore.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
