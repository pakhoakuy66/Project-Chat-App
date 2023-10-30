package routes

import (
	"github.com/gin-gonic/gin"

	"server/controllers"
)

func InitAuthRoute(r *gin.Engine) {
	authRoute := r.Group("/auth")
	authRoute.POST("/register", controllers.Register)
	authRoute.POST("/login", controllers.Login)
	authRoute.POST("/logout", controllers.Logout)
}
