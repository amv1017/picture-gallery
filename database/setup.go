package database

import (
	"gorm.io/gorm"
	"github.com/amv1017/picture-gallery/models"
)

var (

	setup_paintings = []models.Painting {
		{ Title: "Flower", Description: "123" },
		{ Title: "Mountain", Description: "456" },
		{ Title: "Autumn", Description: "789" },
		{ Title: "Tree", Description: "614" },
		{ Title: "Waterfall", Description: "928" },
	}

	setup_author = []models.Author { 
		{ Name: "Levitan", Paintings: []models.Painting{ 
				setup_paintings[0],
				setup_paintings[2],
				setup_paintings[3],
			},  
		},
		{ Name: "Tizian", Paintings: []models.Painting{ 
				setup_paintings[1],
				setup_paintings[4],
			},  
		},
	}

	setup_genre = []models.Genre {
		{ Sign: "Port", Paintings: []models.Painting{ 
				setup_paintings[3],
				setup_paintings[4],
			},
	 	},
		{ Sign: "Land", Paintings: []models.Painting{ 
				setup_paintings[2],
				setup_paintings[0],
			},
	 	},
		{ Sign: "Natu", Paintings: []models.Painting{ 
				setup_paintings[1],
			},
	 	},
	}

)

func FillDatabase(db *gorm.DB) {


	for i := range setup_genre {
		db.Create(&setup_genre[i])
	}

	for i := range setup_author {
		db.Create(&setup_author[i])
	}


}

