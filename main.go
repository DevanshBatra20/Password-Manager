package main

import (
	"github.com/DevanshBatra20/Password-Manager/configs"
	"github.com/DevanshBatra20/Password-Manager/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	configs.ConnectDB()

	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.PasswordRoutes(router)

	router.Run()
}
