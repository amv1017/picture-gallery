package main

import (
	"github.com/amv1017/picture-gallery/controllers"
	"github.com/amv1017/picture-gallery/models"
	"github.com/amv1017/picture-gallery/database"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	database.DB.AutoMigrate(&models.Author{})
	database.DB.AutoMigrate(&models.Painting{})
	database.DB.AutoMigrate(&models.Genre{})

	// WARNING: CALL ONLY THE FIRST TIME
	// database.FillDatabase()

	router := mux.NewRouter()


	router.HandleFunc("/authors", controllers.GetAllAuthors).Methods("GET")
	router.HandleFunc("/author/{id}", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/authors", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/author/{id}", controllers.DeleteAuthor).Methods("DELETE")

	router.HandleFunc("/paintings", controllers.GetAllPaintings).Methods("GET")
	router.HandleFunc("/painting/{id}", controllers.GetPainting).Methods("GET")
	router.HandleFunc("/paintings", controllers.CreatePainting).Methods("POST")
	router.HandleFunc("/painting/{id}", controllers.DeletePainting).Methods("DELETE")

	router.HandleFunc("/genre/{id}", controllers.GetGenre).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client")))


	
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
