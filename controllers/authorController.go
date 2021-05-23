package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/amv1017/picture-gallery/database"
	"github.com/amv1017/picture-gallery/models"
	"github.com/gorilla/mux"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	database.DB.Find(&authors)
	json.NewEncoder(w).Encode(&authors)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author models.Author
	var paintings []models.Painting
	database.DB.First(&author, params["id"])

	author.Paintings = paintings
	json.NewEncoder(w).Encode(&author)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)
	result := database.DB.Create(&author)
	err := result.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&author)
	}
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author models.Author
	database.DB.First(&author, params["id"])
	database.DB.Delete(&author)
	json.NewEncoder(w).Encode(&author)
}
