package controllers

import (
	"bike_manage/config"
	model "bike_manage/models"
	"github.com/gin-gonic/gin"
)

func ShowUserList(c *gin.Context) {
	userList := []model.RegisterUser{}
	if err := config.DB.Find(&userList).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": userList})
}

func CreateNewUser(c *gin.Context) {
	user := model.RegisterUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": user})
}

func ShowBillList(c *gin.Context) {
	billList := []model.UserBill{}
	if err := config.DB.Find(&billList).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": billList})
}
func CreateBill(c *gin.Context) {
	bill := model.UserBill{}
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	if err := config.DB.Create(&bill).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": bill})
}

func CreateFeedback(c *gin.Context) {
	feedback := model.Feedback{}
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	//判断用户id是否存在
	var user model.RegisterUser
	if err := config.DB.Where("id = ?", feedback.UserID).First(&user).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	if err := config.DB.Create(&feedback).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": feedback})
}

func FeedbackList(c *gin.Context) {
	var feedbackList []model.Feedback
	if err := config.DB.Find(&feedbackList).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": feedbackList})
}
