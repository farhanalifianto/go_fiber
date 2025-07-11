package models

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"-"`
	Role      string    `json:"role"` // user / admin
	Notes     []Note    `json:"notes"`
	Favorites []Note    `gorm:"many2many:user_favorites" json:"favorites"`
	Products  []Product `json:"products"`
}