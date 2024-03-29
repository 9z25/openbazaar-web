package main

import (
	"log"
	"net/http"

	"github.com/9z25/openbazaar-web/app"
	"github.com/9z25/openbazaar-web/db"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   r,
		Database: database,
	}

	app.SetupRouter()

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
