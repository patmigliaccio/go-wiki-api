package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	apiv1 := router.PathPrefix("/api/v1.0").Subrouter()
	apiv1.HandleFunc("/extracts/{titles}", getExtracts).Methods("GET")
	apiv1.HandleFunc("/search/{value}", getSearch).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,
		`Welcome to the Wikipedia API.

Please use one of the following endpoints: 
GET /api/v1.0/extracts/{titles}
GET /api/v1.0/search/{value}`)
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
