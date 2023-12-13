package main

import (
	"github.com/DevanshBatra20-PasswordManager/datastore"
	"github.com/DevanshBatra20-PasswordManager/handlers"
	"github.com/DevanshBatra20-PasswordManager/middlewares"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	userDatastore := datastore.NewUser()
	userHandler := handlers.NewUser(userDatastore)

	authDatastore := datastore.NewAuth()
	authHandler := handlers.NewAuth(authDatastore)

	app.POST("/users/signup", authHandler.Signup)
	app.POST("/users/login", authHandler.Login)
	app.GET("/users/{userId}", middlewares.JWTAuth(userHandler.GetById))
	app.Start()
}
