package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null`
	Password string `gorm:"not null"`
}
