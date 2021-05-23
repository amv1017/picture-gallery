package controllers

import (
	"net/http"
	"github.com/amv1017/picture-gallery/database"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
)

type Result struct {
	PaintingID		int		`json:"id"`
	PaintingTitle	string	`json:"title"`
	PaintingUrl		string	`json:"url"`
}

func GetGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	
	// database.DB.Table("genre_paintings").Where("genre_id",params["id"]).Find(&paintings)


	var result []Result
	fmt.Println(params["id"])
	database.DB.Raw("SELECT painting_id, painting_title, painting_url FROM genre_paintings WHERE genre_id = "+params["id"]).Scan(&result)

	fmt.Println(result)

	json.NewEncoder(w).Encode(&result)
}
