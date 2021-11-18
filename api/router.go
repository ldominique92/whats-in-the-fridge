package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleHttpRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/recipes", ListRecipes).Methods("GET").Name("list_recipes")
	log.Fatal(http.ListenAndServe(":10000", router))
}
