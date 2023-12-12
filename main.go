package main

import (
	"github.com/DevanshBatra20-PasswordManager/handlers"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	handlers.AuthHandler(app)

	app.Start()
}
