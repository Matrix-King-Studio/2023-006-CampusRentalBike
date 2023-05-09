package controllers

import (
	"bike_manage/config"
	model "bike_manage/models"
	"github.com/gin-gonic/gin"
)

func SetPrice(c *gin.Context) {
	price := model.BikePrice{}
	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	if err := config.DB.Create(&price); err.Error != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": price})
}

func PriceList(c *gin.Context) {
	var priceList []model.BikePrice
	if err := config.DB.Find(&priceList).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": priceList})
}
func ChangePrice(c *gin.Context) {
	price := model.BikePrice{}
	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	test := model.BikePrice{}
	// 检查是否存在
	if err := config.DB.Where("id = ?", price.ID).First(&test).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	if err := config.DB.Model(&price).Where("id = ?", price.ID).Update(map[string]interface{}{"price": price.Price, "start_price": price.StartPrice, "km_price": price.KmPrice, "deposit": price.Deposit}).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "detail": price})
}

func DeletePrice(c *gin.Context) {
	price := model.BikePrice{}
	ID := c.Query("id")
	test := model.BikePrice{}
	// 检查是否存在
	if err := config.DB.Where("id = ?", ID).First(&test).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	if err := config.DB.Where("id = ?", ID).Delete(&price).Error; err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}
