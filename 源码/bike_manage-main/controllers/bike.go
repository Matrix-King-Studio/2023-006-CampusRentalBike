package controllers

import (
	"bike_manage/config"
	model "bike_manage/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"net/http"
	"strconv"
)

type BikeRequest struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	BelongFence []uint  `json:"belong_fence"`
}

type BikeRecordRequest struct {
	BikeID    uint    `json:"bike_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// CreateBike create bike
func CreateBike(c *gin.Context) {
	var bikeRequest BikeRequest
	if err := c.BindJSON(&bikeRequest); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to parse request body",
		})
		return
	}

	location := model.Location{
		Latitude:  bikeRequest.Latitude,
		Longitude: bikeRequest.Longitude,
	}

	// Save the location
	//config.DB.Create(&location)

	bike := model.Bike{
		IsFault: false,
		IsOut:   false,
	}
	bike.Locations = append(bike.Locations, location)
	// 设置归属限行区
	for _, fId := range bikeRequest.BelongFence {
		var fences model.Fence
		if err := config.DB.Where("id = (?)", fId).Find(&fences).Error; err != nil {
			c.JSON(404, gin.H{"status": "fail"})
			return
		}
		bike.BelongFence = append(bike.BelongFence, fences)
	}
	config.DB.Create(&bike)

	c.JSON(http.StatusOK, gin.H{"status": "success", "detail": gin.H{"bike_id": bike.ID}})
}

func GetHistoryLocation(c *gin.Context) {
	var bike model.Bike
	id := c.Query("bike_id")
	if err := config.DB.Preload("Locations").Where("id = (?)", id).First(&bike).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bike ID"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "details": bike.Locations})
}

// RecordLocation record
func RecordLocation(c *gin.Context) {
	var bikeRecord BikeRecordRequest
	if err := c.ShouldBindJSON(&bikeRecord); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to parse request body",
		})
		return
	}
	location := model.Location{
		Latitude:  bikeRecord.Latitude,
		Longitude: bikeRecord.Longitude,
	}
	var bike model.Bike
	if err := config.DB.Preload("BelongFence").Where("id = (?)", bikeRecord.BikeID).First(&bike).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bike ID"})
		return
	}
	bike.Locations = append(bike.Locations, location)

	for _, f := range bike.BelongFence {
		if !f.IsPointInside(location.Latitude, location.Longitude) {
			bike.IsOut = true
		}
	}

	config.DB.Save(&bike)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func BikePages(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var bikes []model.Bike
	//result := config.DB.Scopes(Paginate(page, limit)).Find(&bikes)
	result := config.DB.Scopes(Paginate(page, limit)).Preload("BelongFence").Find(&bikes)
	// 预加载每个Bike的最新Location
	for i := range bikes {
		var lastLocation model.Location
		config.DB.Where("bike_id = ?", bikes[i].ID).Last(&lastLocation)
		bikes[i].Locations = []model.Location{lastLocation}
	}

	// 查询Bike总数
	var totalBikes int64
	config.DB.Model(&model.Bike{}).Count(&totalBikes)

	// 计算总页数
	totalPages := int(math.Ceil(float64(totalBikes) / float64(limit)))

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		// 返回Bike列表和总页数
		c.JSON(200, gin.H{
			"bikes":       bikes,
			"total_pages": totalPages,
		})
	}
}

// Paginate 创建分页作用域
func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

// Summary 首页统计信息
func Summary(c *gin.Context) {
	var bikes model.Bike
	var TotalBike int64
	var FaultBike int64
	var OutBike int64
	if err := config.DB.Model(&bikes).Count(&TotalBike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	if err := config.DB.Model(&bikes).Where("is_fault = (?)", true).Count(&FaultBike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	if err := config.DB.Model(&bikes).Where("is_out = (?)", true).Count(&OutBike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "success", "detail": gin.H{"TotalBikes": TotalBike, "FaultBikes": FaultBike, "OutBikes": OutBike}})
}

func AllBikeLocation(c *gin.Context) {
	var bikes []model.Bike
	if err := config.DB.Preload("Locations").Find(&bikes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	type ReStruct struct {
		ID        uint
		Locations []model.Location
	}
	var reStruct []ReStruct
	for _, v := range bikes {
		reStruct = append(reStruct, ReStruct{ID: v.ID, Locations: v.Locations})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "detail": reStruct})
}

func DeleteBike(c *gin.Context) {
	var bike model.Bike
	id := c.Query("bike_id")
	if err := config.DB.Where("id = (?)", id).Delete(&bike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func SearchBike(c *gin.Context) {
	var bike model.Bike
	id := c.Query("bike_id")
	if err := config.DB.Preload("Locations").Where("id = (?)", id).First(&bike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "detail": bike})
}
