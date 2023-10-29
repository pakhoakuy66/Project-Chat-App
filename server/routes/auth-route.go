package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"server/controllers"
)

func InitAuthRoute(r *gin.Engine) {
	authRoute := r.Group("/auth")
	authRoute.GET("/test", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("Authorization")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(cookie)
	})
	authRoute.POST("/register", controllers.Register)
	authRoute.POST("/login", controllers.Login)
	authRoute.POST("/logout", controllers.Logout)
}
