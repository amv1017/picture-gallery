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

	fmt.Println(painting)

	var genre models.Genre
	
	database.DB.Model(&models.Genre{}).First(&genre,"sign",painting.Genre)

	fmt.Println(genre)

	database.DB.Exec(`INSERT INTO public.genre_paintings (genre_id,genre_sign,painting_id,painting_title) VALUES (`+
	strconv.Itoa(int(genre.ID))+`,'`+
	genre.Sign+`',`+strconv.Itoa(int(painting.ID))+`,'`+painting.Title+`')`)
	

}

func DeletePainting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var painting models.Painting
	database.DB.First(&painting, params["id"])
	database.DB.Delete(&painting)
	json.NewEncoder(w).Encode(&painting)
}

