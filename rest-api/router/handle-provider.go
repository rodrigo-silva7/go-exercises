package router

import (
	"alura-rest-api/controller"
   "alura-rest-api/middleware"
	"net/http"

   "github.com/gorilla/mux"
)

var (
   PORT = "8080"
)

func HandleProvider() {
   r := mux.NewRouter()
   r.Use(middleware.ContentType)

   r.HandleFunc("/", controller.Home)
   r.HandleFunc("/api/persons", controller.AllPersons).Methods("Get")
   r.HandleFunc("/api/persons", controller.NewPerson).Methods("Post")
   r.HandleFunc("/api/persons/{id}", controller.GetPerson).Methods("Get")
   r.HandleFunc("/api/persons/{id}", controller.DeletePerson).Methods("Delete")
   r.HandleFunc("/api/persons/{id}", controller.EditPerson).Methods("Put")

   http.ListenAndServe(":"+PORT, r)
}
