package models

import (
	"gorm.io/gorm"
)

type Painting struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
}
