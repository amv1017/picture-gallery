package models

type Genre struct {
	ID		uint	`json:"id" gorm:"primary_key"`
	GenreId uint	`json:"genreid"`
	Label	string	`json:"label"`
}

