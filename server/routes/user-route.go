package routes

import (
	"github.com/gin-gonic/gin"

	"server/controllers"
)

func InitUserRoute(r *gin.Engine) {
	userRoute := r.Group("/user")
	userRoute.POST("/register", controllers.Register)
	userRoute.POST("/login", controllers.Login)
}
