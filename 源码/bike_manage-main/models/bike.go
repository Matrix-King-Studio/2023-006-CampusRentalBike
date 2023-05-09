package models

import (
	"github.com/jinzhu/gorm"
)

type Bike struct {
	gorm.Model
	IsFault     bool
	Locations   []Location
	IsOut       bool
	BelongFence []Fence    `gorm:"many2many:fence_bikes;"`
	Invoices    []Invoices `gorm:"foreignKey:UserID"`
}
