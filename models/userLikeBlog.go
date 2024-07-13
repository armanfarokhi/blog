package models

import "github.com/jinzhu/gorm"

type UserLikeBlog struct {
	gorm.Model
	UserID uint
	BlogID uint
	User   User
	Blog   Blog
}
