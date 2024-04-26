package routes

import (
	"github.com/DevanshBatra20/Password-Manager/controllers"

	"github.com/DevanshBatra20/Password-Manager/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.JWTAuth())
	user := incomingRoutes.Group("api/v1/")
	{
		user.GET("/user/:user_id", controllers.GetUser())
		user.POST("/user/uploadImage/:user_id", controllers.UploadImage())
	}
}
