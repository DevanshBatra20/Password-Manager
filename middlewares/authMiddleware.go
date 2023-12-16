package middlewares

import (
	"github.com/DevanshBatra20-PasswordManager/helpers"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func JWTAuth(next gofr.Handler) gofr.Handler {
	return func(ctx *gofr.Context) (interface{}, error) {
		token := helpers.ExtractToken(ctx.Request())
		_, err := helpers.ValidateToken(token)
		if err != nil {
			return nil, &errors.Response{
				StatusCode: 401,
				Reason:     "Invalid Token",
			}
		}

		return next(ctx)
	}
}
