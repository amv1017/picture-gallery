package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/amv1017/picture-gallery/controllers"
	"github.com/amv1017/picture-gallery/database"
	"github.com/amv1017/picture-gallery/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	var dotEnv map[string]string
	dotEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Failed to get environment variables")
	}

	pgURL, err := pq.ParseURL(dotEnv["ELEPHANTSQL_URL"])
	database.DB, err = gorm.Open(postgres.Open(pgURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	database.DB.AutoMigrate(&models.Author{})
	database.DB.AutoMigrate(&models.Painting{})
	database.DB.AutoMigrate(&models.Genre{})

	// WARNING: CALL ONLY THE FIRST TIME WITH EMPTY DATABASE
	// database.FillDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/api/authors", controllers.GetAllAuthors).Methods("GET")
	router.HandleFunc("/api/author/{id}", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/api/authors", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/api/author/{id}", controllers.DeleteAuthor).Methods("DELETE")

	router.HandleFunc("/api/paintings", controllers.GetAllPaintings).Methods("GET")
	router.HandleFunc("/api/painting/{id}", controllers.GetPainting).Methods("GET")
	router.HandleFunc("/api/paintings", controllers.CreatePainting).Methods("POST")
	router.HandleFunc("/api/painting/{id}", controllers.DeletePainting).Methods("DELETE")

	router.HandleFunc("/api/genre/{id}", controllers.GetGenre).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client")))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
