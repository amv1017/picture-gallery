package database

import (
	"gorm.io/gorm"
	"github.com/amv1017/picture-gallery/models"
)

var (

	setup_paintings = []models.Painting {
		{ Title: "A", Description: "123",
			// Genre: setup_genre[1]
			Genre: setup_genre[2],
		},
		{ Title: "B", Description: "456",
			// Genre: setup_genre[0]
		},
		{ Title: "C", Description: "789",
			// Genre: setup_genre[3]
		},
	}

	setup_author = []models.Author { 
		{ Name: "Bruegel", Paintings: setup_paintings[0:2] },
		{ Name: "Bosch", Paintings: setup_paintings[2:] },
	}

	setup_genre = []models.Genre {
		{ Label: "Port" },
		{ Label: "Natu" },
		{ Label: "Livi" },
		{ Label: "Hist" },
	}

)

func FillDatabase(db *gorm.DB) {


	for i := range setup_author {
		db.Create(&setup_author[i])
	}


}

