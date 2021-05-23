package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/amv1017/picture-gallery/database"
	"github.com/amv1017/picture-gallery/models"
	"github.com/gorilla/mux"
)


func GetAllAuthors(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var authors []models.Author
	database.DB.Find(&authors)
	json.NewEncoder(w).Encode(&authors)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	var result []Result
	database.DB.Raw("SELECT painting_id, painting_title, painting_url FROM author_paintings WHERE author_id = "+params["id"]).Scan(&result)
	json.NewEncoder(w).Encode(&result)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	var author models.Author
	database.DB.First(&author, params["id"])

	database.DB.Exec(`DELETE FROM public.author_paintings WHERE author_id=`+
	strconv.Itoa(int(author.ID))+` AND author_name='`+author.Name+`'`)

	database.DB.Exec(`DELETE FROM public.authors WHERE id=`+
	strconv.Itoa(int(author.ID))+` AND name='`+author.Name+`'`)

	json.NewEncoder(w).Encode(&author)
}
