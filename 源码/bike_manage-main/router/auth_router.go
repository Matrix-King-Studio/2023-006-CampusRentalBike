package router

import (
	"bike_manage/controllers"
	"github.com/gin-gonic/gin"
)

func InitAuth(router *gin.Engine) {
	auth := router.Group("auth/")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/wxlogin", controllers.WeChatLogin)
	}
	return
}
