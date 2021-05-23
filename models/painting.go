package models

type Painting struct {
	ID			uint	`json:"id" gorm:"primary_key;unique"`
	Title		string	`json:"title" gorm:"primary_key"`
	Description	string	`json:"description"`
	Url			string	`json:"url"`
	Genre		string	`json:"genre"`
	Author		string	`json:"author"`
}
