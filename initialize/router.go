package initialize

import (
	"contact/middleware"
	contactRouter "contact/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	apiRouter := router.Group("")
	contactRouter.InitContactRouter(apiRouter)

	return router
}
