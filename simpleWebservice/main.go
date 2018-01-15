package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/skrelan/GOdev/simpleWebservice/middleware"
)

func endpoints(router *mux.Router) {
	router.HandleFunc("/people", middleware.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", middleware.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", middleware.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", middleware.DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/people/{id}", middleware.UpdatePersonEndpoint).Methods("PUT")
}

func main() {
	middleware.Start()
	router := mux.NewRouter()
	endpoints(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
