package initialize

import (
	"contact/middleware"
	contactRouter "contact/router"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	apiRouter := router.Group("")
	contactRouter.InitContactRouter(apiRouter)

	return router
}
