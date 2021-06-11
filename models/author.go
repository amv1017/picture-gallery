package models

type Author struct {
	ID			uint		`json:"id" gorm:"primary_key"`
	Name		string		`json:"name" gorm:"primary_key"`
	Paintings	[]Painting	`json:"paintings" gorm:"many2many:author_paintings"`
}
