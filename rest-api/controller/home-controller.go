package controller

import (
	"alura-rest-api/database"
	"alura-rest-api/model"
	"encoding/json"
	"fmt"
   "strconv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w, "Home Page")
}

func AllPersons(w http.ResponseWriter, r*http.Request) {
   var p []model.Person
   database.DB.Find(&p)

   json.NewEncoder(w).Encode(p)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   id := vars["id"]
   var person model.Person
   database.DB.First(&person, id)
   json.NewEncoder(w).Encode(person)
}

func NewPerson(w http.ResponseWriter, r *http.Request) {
   var person model.Person
   json.NewDecoder(r.Body).Decode(&person)
   database.DB.Create(&person)
   json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   id := vars["id"]
   var person model.Person
   database.DB.First(&person, id)
   log.Print("Deletando:", person)
   database.DB.Delete(&person, id)
   json.NewEncoder(w).Encode(person)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   id := vars["id"]
   var person model.Person
   database.DB.First(&person, id)
   log.Print("Editando: ", person)

   json.NewDecoder(r.Body).Decode(&person)
   person.Id, _ = strconv.Atoi(id)
   database.DB.Save(&person)

   json.NewEncoder(w).Encode(person)
}

