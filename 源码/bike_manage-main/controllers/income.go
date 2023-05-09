package controllers

import (
	"bike_manage/config"
	model "bike_manage/models"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"time"
)

type RequestIncome struct {
	Second int  `json:"second"`
	UserID uint `json:"user_id"`
	BikeID uint `json:"bike_id"`
}

func calculateFee(seconds int, x int, y float64, t float64) float64 {
	// 将起步时间转换为秒
	initialSeconds := x * 60
	// 计算起步费用
	if seconds <= initialSeconds {
		return y
	}

	// 计算超过起步时间的秒数
	remainingSeconds := seconds - initialSeconds
	// 将剩余秒数转换为 30 分钟的倍数，向上取整
	numThirtyMinuteBlocks := math.Ceil(float64(remainingSeconds) / 1800)

	// 计算总费用
	totalFee := y + (numThirtyMinuteBlocks * t)
	return totalFee
}
func MakeIncome(c *gin.Context) {
	var RIncome RequestIncome
	var income model.Invoices
	if err := c.BindJSON(&RIncome); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fee := calculateFee(RIncome.Second, 30, 1, 1.5)
	income.BikeID = RIncome.BikeID
	income.UserID = RIncome.UserID
	income.Amount = fee
	config.DB.Save(&income)
	c.JSON(200, gin.H{"status": "success", "detail": income})
}

// HistoryDayIncome 返回以天为单位，返回数据库中每天的总收益,并返回某一天的信息

func HistoryDayIncome(c *gin.Context) {
	var invoices []model.Invoices
	config.DB.Table("invoices").Select("DATE(created_at) as created_at, SUM(amount) as amount").Group("DATE(created_at)").Scan(&invoices)
	c.JSON(200, gin.H{"status": "success", "detail": invoices})
}
func HistoryBill(c *gin.Context) {
	var bills []model.Invoices
	config.DB.Preload("User").Find(&bills)

	// Create a list of anonymous structs with an additional field for the username
	var responseBills []struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Amount    float64   `json:"amount"`
		UserID    uint      `json:"user_id"`
		BikeID    uint      `json:"bike_id"`
		Username  string    `json:"username"`
	}

	for _, bill := range bills {
		responseBills = append(responseBills, struct {
			ID        uint      `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			Amount    float64   `json:"amount"`
			UserID    uint      `json:"user_id"`
			BikeID    uint      `json:"bike_id"`
			Username  string    `json:"username"`
		}{
			ID:        bill.ID,
			CreatedAt: bill.CreatedAt,
			Amount:    bill.Amount,
			UserID:    bill.UserID,
			BikeID:    bill.BikeID,
			Username:  bill.User.Username,
		})
	}

	c.JSON(200, gin.H{"status": "success", "detail": responseBills})
}

// IncomeInformation 返回总收益，损坏单车数量，总单车数量
func IncomeInformation(c *gin.Context) {
	var invoices []model.Invoices
	var bikes []model.Bike
	var total float64
	var totalFault int
	config.DB.Table("invoices").Select("SUM(amount) as amount").Scan(&invoices)
	config.DB.Find(&bikes)
	for _, v := range invoices {
		total += v.Amount
	}
	for _, v := range bikes {
		if v.IsFault {
			totalFault += 1
		}
	}
	totalBikeAmount := len(bikes)
	totalBikeAmount *= 150

	totalFaultAmount := totalFault * 150

	c.JSON(200, gin.H{"status": "success", "total": total, "totalBikeAmount": totalBikeAmount, "totalFaultAmount": totalFaultAmount})
}
