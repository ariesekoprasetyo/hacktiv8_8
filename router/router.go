package router

import (
	"Assigment_8/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	routerGroup := router.Group("/api/v1")
	{
		routerGroup.PUT("/status", controllers.StatusUpdate)
	}
	return
}
