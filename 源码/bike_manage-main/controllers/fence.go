package controllers

import (
	"bike_manage/config"
	"bike_manage/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateFence(c *gin.Context) {
	var fence models.Fence
	if err := c.ShouldBindJSON(&fence); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	config.DB.Save(&fence)
	c.JSON(200, gin.H{"status": "success", "detail": fence})
}

func AllFence(c *gin.Context) {
	var fence []models.Fence
	type PureLocation struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	type reStruct struct {
		Name      string         `json:"name"`
		Id        uint           `json:"id"`
		CreatedAt time.Time      `json:"createdAt"`
		UpdatedAt time.Time      `json:"updatedAt"`
		Locations []PureLocation `json:"locations"`
	}

	var reSlice []reStruct
	// 预先加载locations
	if err := config.DB.Preload("Locations").Find(&fence).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	for _, f := range fence {
		re := reStruct{Id: f.ID, Name: f.Name, CreatedAt: f.CreatedAt, UpdatedAt: f.UpdatedAt}
		for _, l := range f.Locations {
			re.Locations = append(re.Locations, PureLocation{Latitude: l.Latitude, Longitude: l.Longitude})
		}
		reSlice = append(reSlice, re)
	}

	c.JSON(200, gin.H{"status": "success", "detail": reSlice})
}

func DeleteFence(c *gin.Context) {
	fenceId := c.Param("id")
	var fence models.Fence
	if err := config.DB.Where("id = ?", fenceId).First(&fence).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	config.DB.Delete(&fence)
	c.JSON(200, gin.H{"status": "success"})
}
