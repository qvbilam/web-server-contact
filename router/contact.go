package router

import (
	"contact/api/conversation"
	"contact/api/friend"
	"contact/api/group"
	"contact/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitContactRouter(Router *gin.RouterGroup) {
	ContactRouter := Router.Group("/contact/").Use(middleware.Cors()).Use(middleware.Auth())
	{
		ContactRouter.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "pong",
			})
		})

		ContactRouter.GET("version", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "v1.0.0",
			})
		})

		ContactRouter.GET("friend", friend.Friends)       // 好友列表
		ContactRouter.PUT("friend/:id", friend.Update)    // 修改好友信息
		ContactRouter.DELETE("friend/:id", friend.Delete) // 删除好友

		ContactRouter.GET("friend/apply", friend.News)          // 申请列表
		ContactRouter.POST("friend/apply", friend.Apply)        // 申请好友
		ContactRouter.PUT("friend/apply/:id", friend.Agree)     // 同意申请
		ContactRouter.DELETE("friend/apply/:id", friend.Reject) // 拒绝申请

		ContactRouter.POST("group", group.Create)            // 创建群
		ContactRouter.GET("group", group.Search)             // 搜索群
		ContactRouter.GET("group/mine", group.Mine)          // 我的群
		ContactRouter.GET("group/:id/member", group.Members) // 群成员
		ContactRouter.POST("group/:id/member", group.Join)   // 加入群
		ContactRouter.DELETE("group/:id/member", group.Quit) // 退出群

		ContactRouter.GET("conversation", conversation.Get)           // 获取最近联系人
		ContactRouter.POST("conversation", conversation.Create)       // 创建联系人
		ContactRouter.DELETE("conversation/:id", conversation.Delete) // 删除联系人
	}
}
