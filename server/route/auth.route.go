package route

import (
	"github.com/gin-gonic/gin"

	"server/controller"
)

func InitAuthRoute(r *gin.Engine) {
	authRoute := r.Group("/auth")
	authRoute.POST("/register", controller.Register)
	authRoute.POST("/login", controller.Login)
	authRoute.POST("/refresh", controller.Refresh)
}
