package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Response is the default object returned from the API
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	// Allows cross-origin requests from any domain
	o := handlers.AllowedOrigins([]string{"*"})

	apiv1 := router.PathPrefix("/api/v1.0").Subrouter()
	apiv1.HandleFunc("/extracts/{titles}", GetExtracts).Methods("GET")
	apiv1.HandleFunc("/search/{value}", GetSearch).Queries("limit", "{limit}").Methods("GET")
	apiv1.HandleFunc("/search/{value}", GetSearch).Methods("GET")
	apiv1.HandleFunc("/categories/{pageid}", GetCategories).Methods("GET")
	apiv1.HandleFunc("/sections/{pageid}", GetSections).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(o)(router)))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,
		`Welcome to the Wikipedia API.

Please use one of the following endpoints: 
GET /api/v1.0/extracts/{titles}
GET /api/v1.0/search/{value}?limit={limit}
GET /api/v1.0/categories/{pageid}
GET /api/v1.0/sections/{pageid}`)
}
