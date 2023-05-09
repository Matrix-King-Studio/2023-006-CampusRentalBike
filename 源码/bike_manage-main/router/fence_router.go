package router

import (
	"bike_manage/controllers"
	"bike_manage/middlewares"
	"github.com/gin-gonic/gin"
)

func InitFenceRouter(router *gin.Engine) {
	Fence := router.Group("/fence")
	Fence.Use(middlewares.JWTAuth())
	{
		Fence.POST("/create", controllers.CreateFence)
		Fence.GET("/list", controllers.AllFence)
		Fence.DELETE("/delete/:id", controllers.DeleteFence)
	}
}
