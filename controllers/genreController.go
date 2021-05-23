package controllers

import (
	"net/http"
	"github.com/amv1017/picture-gallery/database"
	"github.com/gorilla/mux"
	"encoding/json"
)

func GetGenre(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	var result []Result
	database.DB.Raw("SELECT painting_id, painting_title, painting_url FROM genre_paintings WHERE genre_id = "+params["id"]).Scan(&result)
	json.NewEncoder(w).Encode(&result)
}
