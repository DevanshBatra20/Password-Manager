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

	passwordDatastore := datastore.NewPassword()
	passwordHandler := handlers.NewPassword(passwordDatastore)

	app.POST("/users/signup", authHandler.Signup)
	app.POST("/users/login", authHandler.Login)
	app.GET("/users/{userId}", middlewares.JWTAuth(userHandler.GetById))
	app.DELETE("/users/deleteUser/{userId}", middlewares.JWTAuth(userHandler.DeleteById))
	app.POST("/users/createPassword/{userId}", middlewares.JWTAuth(passwordHandler.Create))
	app.DELETE("/users/deletePassword/{passwordId}", middlewares.JWTAuth(passwordHandler.Delete))
	app.PUT("/users/updatePassword/{passwordId}", middlewares.JWTAuth(passwordHandler.Update))
	app.GET("/users/getPassword/{passwordId}", middlewares.JWTAuth(passwordHandler.Get))
	app.GET("/users/getAllPasswords/{userId}", middlewares.JWTAuth(passwordHandler.GetByUserId))

	app.Start()
}
