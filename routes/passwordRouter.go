package routes

import (
	"github.com/DevanshBatra20/Password-Manager/controllers"
	"github.com/DevanshBatra20/Password-Manager/middleware"
	"github.com/gin-gonic/gin"
)

func PasswordRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.JWTAuth())
	password := incomingRoutes.Group("/api/v1")
	{
		password.POST("/password/:user_id", controllers.CreatePassword())
	}
}
