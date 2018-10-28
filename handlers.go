package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetExtracts returns the content for a list of titles
func GetExtracts(w http.ResponseWriter, r *http.Request) {
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

// GetSearch retrieves a list of pages based on a search value.
func GetSearch(w http.ResponseWriter, r *http.Request) {
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

// GetCategories retrieves the categories associated with an article.
func GetCategories(w http.ResponseWriter, r *http.Request) {
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

// GetSections retrieves the sections within an article.
func GetSections(w http.ResponseWriter, r *http.Request) {
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
