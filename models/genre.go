package models

type Genre struct {
	ID		uint	`json:"id" gorm:"primary_key"`
	Sign	string	`json:"description" gorm:"primary_key"`
	Paintings	[]Painting	`json:"paintings" gorm:"many2many:genre_paintings"`
}

