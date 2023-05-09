package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name        string        `gorm:"unique_index"`
	Permissions []*Permission `gorm:"many2many:role_permissions;"`
}
