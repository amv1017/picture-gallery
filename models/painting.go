package models

type Painting struct {
	ID			uint	`json:"id" gorm:"primary_key;unique"`
	Title		string	`json:"title" gorm:"primary_key"`
	Description	string	`json:"description"`
	Url			string	`json:"url" gorm:"primary_key"`
	Author		string	`json:"author"`
	Genre		string	`json:"genre"`
}
