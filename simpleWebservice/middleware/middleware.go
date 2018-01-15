package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skrelan/GOdev/simpleWebservice/models"
	"github.com/skrelan/GOdev/simpleWebservice/utils"
	"github.com/skrelan/LogrusWrapper/log"
)

var people []models.Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Person Not found")
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person models.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			json.NewEncoder(w).Encode("Person Deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Person Not found")
}

func UpdatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person models.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
}

func Start() {
	log.Info("created People")
	utils.CreateData(&people)
}
