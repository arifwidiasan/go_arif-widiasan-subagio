package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Pages  int    `json:"pages" form:"pages"`
}
