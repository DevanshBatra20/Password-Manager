package handlers

import (
	"github.com/DevanshBatra20-PasswordManager/datastore"
	"github.com/DevanshBatra20-PasswordManager/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func AuthHandler(incomingRequests *gofr.Gofr) {
	incomingRequests.POST("/api/users/signup", func(c *gofr.Context) (interface{}, error) {
		var user models.User

		if err := c.Bind(&user); err != nil {
			c.Logger.Errorf("Error in binding user: %v", err)
			return nil, errors.InvalidParam{Param: []string{"user"}}
		}
		resp, err := datastore.Signup(c, &user)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	incomingRequests.POST("/api/users/login", func(c *gofr.Context) (interface{}, error) {
		var userCredentials models.Login

		if err := c.Bind(&userCredentials); err != nil {
			c.Logger.Errorf("Error in binding user: %v", err)
			return nil, errors.InvalidParam{Param: []string{"user"}}
		}
		resp, err := datastore.Login(c, &userCredentials)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
}
