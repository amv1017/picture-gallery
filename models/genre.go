package models

type Genre struct {
	ID			uint		`json:"id" gorm:"primary_key"`
	Title		string		`json:"name" gorm:"primary_key"`
}

