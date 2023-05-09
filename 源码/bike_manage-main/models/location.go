package models

type Location struct {
	Latitude  float64 `gorm:"type:decimal(9,6)" json:"latitude"`
	Longitude float64 `gorm:"type:decimal(9,6)" json:"longitude"`
	BikeID    uint
	FenceID   uint
}
