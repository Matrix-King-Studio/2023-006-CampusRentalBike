package router

import (
	"bike_manage/controllers"
	"bike_manage/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUser(router *gin.Engine) {
	user := router.Group("/user")
	user.Use(middlewares.JWTAuth())
	//user.Use(middlewares.RequireSuperAdminRole())
	{
		user.GET("/list", controllers.ShowUserList)
		user.POST("/create", controllers.CreateNewUser)
		user.GET("/bill/list", controllers.ShowBillList)
		user.POST("/bill/create", controllers.CreateBill)
		user.POST("/feedback/create", controllers.CreateFeedback)
		user.GET("/feedback/list", controllers.FeedbackList)
	}
}
