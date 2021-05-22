package models

type Painting struct {
	ID			uint	`json:"id" gorm:"primary_key;unique"`
	Title		string	`json:"title" gorm:"primary_key"`
	Description	string	`json:"description"`
	Ref			uint	`json:"ref"`
	Genre		Genre	`json:"genre" gorm:"foreignKey:GenreId"`
	Url			string	`json:"url"`
}
