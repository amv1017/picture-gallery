package database

import (
	"gorm.io/gorm"
	"github.com/amv1017/picture-gallery/models"
)

var (

	setup_paintings = []models.Painting {
		{ Title: "A", Description: "123" },
		{ Title: "B", Description: "456" },
		{ Title: "C", Description: "789" },
	}

	setup_author = []models.Author { 
		{ Name: "Bruegel", Paintings: setup_paintings[0:2] },
		{ Name: "Bosch", Paintings: setup_paintings[2:] },
	}

)

func FillDatabase(db *gorm.DB) {

	for i := range setup_author {
		db.Create(&setup_author[i])
	}
}

