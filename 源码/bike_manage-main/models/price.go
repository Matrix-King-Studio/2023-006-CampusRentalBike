package models

type BikePrice struct {
	ID         uint    `gorm:"primary_key" json:"id"`
	Price      float64 `gorm:"type:decimal(10,2)" json:"price"`
	StartPrice float64 `gorm:"type:decimal(10,2)" json:"start_price"`
	KmPrice    float64 `gorm:"type:decimal(10,2)" json:"km_price"`
	Type_      string  `json:"type"`
	Deposit    float64 `gorm:"type:decimal(10,2)" json:"deposit"`
}
