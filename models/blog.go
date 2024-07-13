package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Content     string `gorm:"not null"`
	AuthorID    uint   `gorm:"not null"`
	AuthorEmail string `gorm:"-"`
	Likes       int    `gorm:"default:0"`
}
