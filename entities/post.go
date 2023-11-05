package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Caption     string
	Image_Url   string
	Date        int64
	ID          string
	Likes       string
	Likes_Count int
	Email       string
}
