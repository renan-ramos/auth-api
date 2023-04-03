package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" binding:"required,email"`
	Password string `biding:"required"`
}
