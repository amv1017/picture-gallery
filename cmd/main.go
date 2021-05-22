package main

import (
	"github.com/amv1017/picture-gallery/models"
	"github.com/amv1017/picture-gallery/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


var db *gorm.DB
var err error

func main() {
	var dotEnv map[string]string
	dotEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Failed to get environment variables")
	}

	dsn :=	"host="+dotEnv["HOST"]+
			" user="+dotEnv["USER"]+
			" password="+dotEnv["PASSWORD"]+
			" dbname="+dotEnv["DBNAME"]+
			" port="+dotEnv["PORT"]+
			" sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Painting{})

	// WARNING: CALL ONLY THE FIRST TIME
	database.FillDatabase(db)

	// Routes
	router := mux.NewRouter()

	router.HandleFunc("/authors", getAllAuthors).Methods("GET")
	router.HandleFunc("/author/{id}", getAuthor).Methods("GET")
	router.HandleFunc("/authors", createAuthor).Methods("POST")
	router.HandleFunc("/author/{id}", deleteAuthor).Methods("DELETE")

	router.HandleFunc("/paintings", getAllPaintings).Methods("GET")
	router.HandleFunc("/painting/{id}", getPainting).Methods("GET")
	router.HandleFunc("/paintings", createPainting).Methods("POST")
	router.HandleFunc("/painting/{id}", deletePainting).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/")))
	log.Fatal(http.ListenAndServe(":8080", router))
}

// --- Authors ---

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	db.Find(&authors)
	json.NewEncoder(w).Encode(&authors)
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author models.Author
	var paintings []models.Painting
	db.First(&author, params["id"])

	author.Paintings = paintings
	json.NewEncoder(w).Encode(&author)
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)
	createdPerson := db.Create(&author)
	err = createdPerson.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&author)
	}
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author models.Author
	db.First(&author, params["id"])
	db.Delete(&author)
	json.NewEncoder(w).Encode(&author)
}

// --- Paintings ---

func getAllPaintings(w http.ResponseWriter, r *http.Request) {
	var paintings []models.Painting
	db.Find(&paintings)
	json.NewEncoder(w).Encode(&paintings)
}

func getPainting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var painting models.Painting
	db.First(&painting, params["id"])
	json.NewEncoder(w).Encode(&painting)
}

func createPainting(w http.ResponseWriter, r *http.Request) {
	var book models.Painting
	json.NewDecoder(r.Body).Decode(&book)
	createdBook := db.Create(&book)
	err = createdBook.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&book)
	}
}

func deletePainting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Painting
	db.First(&book, params["id"])
	db.Delete(&book)
	json.NewEncoder(w).Encode(&book)
}

