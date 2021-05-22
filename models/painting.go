package models

type Painting struct {
	ID			uint	`json:"id" gorm:"primary_key"`
	Title		string	`json:"title" gorm:"primary_key"`
	Description	string	`json:"description"`

	Genre		Genre	`json:"genre" gorm:"foreignKey:Title"`
	Url			string	`json:"url"`
}
