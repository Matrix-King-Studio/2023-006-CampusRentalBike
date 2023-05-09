package router

import (
	"bike_manage/controllers"
	"bike_manage/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(router *gin.Engine) {
	customer := router.Group("/customer")
	customer.Use(middlewares.JWTAuth())
	{
		customer.GET("/info", controllers.CustomerInfo)
	}
	return
}
