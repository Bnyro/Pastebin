package entities

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Title 	string `json:"title"`
	Content string `json:"content"`
	Expires uint   `json:"expires"`
}