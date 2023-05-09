package models

import "github.com/jinzhu/gorm"

type Invoices struct {
	gorm.Model
	Amount float64 `gorm:"type:decimal(10,2)" json:"amount"`
	UserID uint    `json:"user_id"`
	BikeID uint    `json:"bike_id"`
	User   User    `gorm:"foreignKey:UserID"`
}
