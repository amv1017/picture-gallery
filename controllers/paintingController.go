package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/amv1017/picture-gallery/database"
	"github.com/amv1017/picture-gallery/models"
	"github.com/gorilla/mux"
)

func GetAllPaintings(w http.ResponseWriter, r *http.Request) {
	var paintings []models.Painting
	database.DB.Find(&paintings)
	json.NewEncoder(w).Encode(&paintings)
}

func GetPainting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var painting models.Painting
	database.DB.First(&painting, params["id"])
	json.NewEncoder(w).Encode(&painting)
}

func CreatePainting(w http.ResponseWriter, r *http.Request) {
	var painting models.Painting
	json.NewDecoder(r.Body).Decode(&painting)
	result := database.DB.Create(&painting)

	err := result.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&painting)
	}

	var genre models.Genre
	var author models.Author

	database.DB.Model(&models.Genre{}).First(&genre,"sign",painting.Genre)
	database.DB.Model(&models.Author{}).First(&author,"name",painting.Author)

	database.DB.Exec(`INSERT INTO public.genre_paintings (genre_id,genre_sign,painting_id,painting_title) VALUES (`+
	strconv.Itoa(int(genre.ID))+`,'`+genre.Sign+`',`+
	strconv.Itoa(int(painting.ID))+`,'`+painting.Title+`')`)

	database.DB.Exec(`INSERT INTO public.author_paintings (author_id,author_name,painting_id,painting_title) VALUES (`+
	strconv.Itoa(int(author.ID))+`,'`+author.Name+`',`+
	strconv.Itoa(int(painting.ID))+`,'`+painting.Title+`')`)

}

func DeletePainting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var painting models.Painting
	database.DB.First(&painting, params["id"])
	fmt.Println(painting)

	var genre models.Genre
	var author models.Author

	database.DB.Model(&models.Genre{}).First(&genre,"sign",painting.Genre)
	database.DB.Model(&models.Author{}).First(&author,"name",painting.Author)

	database.DB.Exec(`DELETE FROM public.genre_paintings WHERE genre_id=`+
	strconv.Itoa(int(genre.ID))+` AND genre_sign='`+genre.Sign+`' AND painting_id=`+
	strconv.Itoa(int(painting.ID))+` AND painting_title='`+painting.Title+`'`)

	database.DB.Exec(`DELETE FROM public.author_paintings WHERE author_id=`+
	strconv.Itoa(int(author.ID))+` AND author_name='`+author.Name+`' AND painting_id=`+
	strconv.Itoa(int(painting.ID))+` AND painting_title='`+painting.Title+`'`)

	database.DB.Exec(`DELETE FROM public.paintings WHERE id=`+
	strconv.Itoa(int(painting.ID))+` AND title='`+painting.Title+`'`)

}

