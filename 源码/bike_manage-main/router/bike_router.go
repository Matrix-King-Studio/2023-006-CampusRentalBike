package router

import (
	"bike_manage/controllers"
	"bike_manage/middlewares"
	"github.com/gin-gonic/gin"
)

func InitBikeRouter(router *gin.Engine) {
	Bike := router.Group("/bike")
	Bike.Use(middlewares.JWTAuth())
	{
		Bike.POST("/create", controllers.CreateBike)
		Bike.GET("/get_location", controllers.GetHistoryLocation)
		Bike.POST("/record_location", controllers.RecordLocation)
		Bike.GET("/bike_lst", controllers.BikePages)
		Bike.GET("/summary", controllers.Summary)
		Bike.GET("/location", controllers.AllBikeLocation)
		Bike.GET("/delete_bike", controllers.DeleteBike)
		Bike.GET("/search_bike", controllers.SearchBike)
	}
}
