package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name string `json:"name"`
	Paintings []Painting `json:"paintings" gorm:"many2many:author_paintings"`
}
