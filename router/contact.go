package router

import (
	"contact/api/friend"
	"contact/middleware"
	"github.com/gin-gonic/gin"
)

func InitContactRouter(Router *gin.RouterGroup) {
	ContactRouter := Router.Group("").Use(middleware.Cors()).Use(middleware.Auth())
	{
		ContactRouter.GET("friend", friend.Friends)
		ContactRouter.PUT("friend/:id", friend.Update)
		ContactRouter.DELETE("friend/:id", friend.Delete)

		ContactRouter.GET("friend/apply", friend.News)
		ContactRouter.POST("friend/apply", friend.Apply)
		ContactRouter.PUT("friend/apply/:id", friend.Agree)
		ContactRouter.DELETE("friend/apply/:id", friend.Reject)

		//ContactRouter.POST("group/publish", api.GroupSend)
	}
}
