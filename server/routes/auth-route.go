package routes

import (
	"github.com/gin-gonic/gin"

	"server/controllers"
	"server/middlewares"
)

func InitAuthRoute(r *gin.Engine) {
	authRoute := r.Group("/auth")
	authRoute.POST("/register",  controllers.Register)
	authRoute.POST("/login",  middlewares.AllowCredentials, controllers.Login)
	authRoute.POST("/logout", middlewares.AllowCredentials, controllers.Logout)
}
