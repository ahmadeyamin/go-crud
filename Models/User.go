package Models

import (
	"gorm.io/gorm"
)

type User struct {
	Name  string `gorm:"name" json:"Name"`
	Email string `gorm:"email" json:"Email"`
	gorm.Model
}
