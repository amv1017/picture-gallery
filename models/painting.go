package models

type Painting struct {
	ID			uint	`json:"id" gorm:"primary_key"`
	Title		string	`json:"title" gorm:"primary_key"`
	Description	string	`json:"description"`
	Url			string	`json:"url"`
}
