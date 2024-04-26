package routes

import (
	"github.com/DevanshBatra20/Password-Manager/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	auth := incomingRoutes.Group("/api/v1/")
	{
		auth.POST("user/signup", controllers.Signup())
		auth.POST("user/login", controllers.Login())
	}
}
