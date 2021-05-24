package controllers

import (
	"net/http"
	"github.com/amv1017/picture-gallery/database"
	"github.com/gorilla/mux"
	"encoding/json"
)

func GetGenre(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var result []Result
	database.DB.Raw("SELECT painting_id, painting_title, painting_url FROM genre_paintings WHERE genre_id = "+params["id"]).Scan(&result)
	json.NewEncoder(w).Encode(&result)

}
