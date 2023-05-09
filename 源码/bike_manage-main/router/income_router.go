package router

import (
	"bike_manage/controllers"
	"bike_manage/middlewares"
	"github.com/gin-gonic/gin"
)

func InitIncomeRouter(router *gin.Engine) {
	Income := router.Group("/income")
	Income.Use(middlewares.JWTAuth())
	{
		Income.POST("/income", controllers.MakeIncome)
		Income.GET("/history", controllers.HistoryDayIncome)
		Income.GET("/bill", controllers.HistoryBill)
		Income.GET("/information", controllers.IncomeInformation)
	}
}
