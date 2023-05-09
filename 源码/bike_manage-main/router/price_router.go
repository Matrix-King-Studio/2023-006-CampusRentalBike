package router

import (
	"bike_manage/controllers"
	"bike_manage/middlewares"
	"github.com/gin-gonic/gin"
)

func InitPriceRouter(router *gin.Engine) {
	Price := router.Group("/price")
	Price.Use(middlewares.JWTAuth())
	{
		Price.POST("/set", controllers.SetPrice)
		Price.GET("/list", controllers.PriceList)
		Price.POST("/change", controllers.ChangePrice)
		Price.GET("/delete", controllers.DeletePrice)
	}
}
