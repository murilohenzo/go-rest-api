package models

import "github.com/jinzhu/gorm"

// Comment - Defines our comment structure
type Comment struct {
	gorm.Model
	Slug string
	Body string
	Author string
}