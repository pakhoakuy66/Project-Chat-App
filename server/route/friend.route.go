package route

import (
	"github.com/gin-gonic/gin"

	"server/controller"
	"server/middleware"
)

func InitFriendRoute(r *gin.Engine) {
	friendRoute := r.Group("friend")
	friendRoute.Use(middleware.Authorize)
	friendRoute.GET("/all", controller.GetAllFriend)
}
