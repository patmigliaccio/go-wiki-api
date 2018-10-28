package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	apiv1.HandleFunc("/extracts/{titles}", getExtracts).Methods("GET")
	apiv1.HandleFunc("/search/{value}", getSearch).Queries("limit", "{limit}").Methods("GET")
	apiv1.HandleFunc("/categories/{pageid}", getCategories).Methods("GET")
	apiv1.HandleFunc("/sections/{pageid}", getSections).Methods("GET")

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

func getExtracts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	wiki, err := NewWikipediaClient()
	if err != nil {
		fmt.Fprintln(w, "Error instantiating Wikipedia Client.")
	}

	titles, err := wiki.GetExtracts([]string{vars["titles"]})
	if err != nil {
		fmt.Fprintln(w, "Error retrieving extracts.")
	}

	json.NewEncoder(w).Encode(titles)
}

func getSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	wiki, err := NewWikipediaClient()
	if err != nil {
		fmt.Fprintln(w, "Error instantiating Wikipedia Client.")
	}

	limit, err := strconv.ParseInt(vars["limit"], 10, 64)
	if err != nil {
		limit = 0
	}

	values, err := wiki.GetPrefixResults(vars["value"], int(limit))
	if err != nil {
		fmt.Fprintln(w, "Error retrieving search values.")
	}

	json.NewEncoder(w).Encode(values)
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	wiki, err := NewWikipediaClient()
	if err != nil {
		fmt.Fprintln(w, "Error instantiating Wikipedia Client.")
	}

	id, err := strconv.ParseInt(vars["pageid"], 10, 64)
	if err != nil {
		id = 0
	}

	value, err := wiki.GetCategories(int(id))
	if err != nil {
		w.WriteHeader(404)

		json.NewEncoder(w).Encode(Response{
			Status:  404,
			Message: "Error retrieving categories.",
		})

		return
	}

	json.NewEncoder(w).Encode(value)
}

func getSections(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	wiki, err := NewWikipediaClient()
	if err != nil {
		fmt.Fprintln(w, "Error instantiating Wikipedia Client.")
	}

	id, err := strconv.ParseInt(vars["pageid"], 10, 64)
	if err != nil {
		id = 0
	}

	value, err := wiki.GetSections(int(id))
	if err != nil {
		w.WriteHeader(404)

		json.NewEncoder(w).Encode(Response{
			Status:  404,
			Message: "Error retrieving sections.",
		})

		return
	}

	json.NewEncoder(w).Encode(value)
}
