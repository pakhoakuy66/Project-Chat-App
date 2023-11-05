package route

import (
	"github.com/gin-gonic/gin"

	"server/controller"
)

func InitFriendRoute(r *gin.Engine) {
	friendRoute := r.Group("friend")
	friendRoute.GET("/:id", controller.GetAllFriend)
}
