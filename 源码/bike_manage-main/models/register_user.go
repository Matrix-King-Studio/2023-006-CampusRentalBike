package models

import "github.com/jinzhu/gorm"

type RegisterUser struct {
	gorm.Model
	Username string     `json:"username"`
	Password string     `json:"password"`
	Bill     []UserBill `json:"bill" gorm:"foreignkey:UserID"`
}

type UserBill struct {
	gorm.Model
	Username  string     `json:"username"`
	Bill      int        `json:"bill"`
	UserID    uint       `json:"user_id"`
	Feedbacks []Feedback `gorm:"foreignkey:UserID"`
}

type Feedback struct {
	gorm.Model
	Username string `json:"username"`
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"`
}
