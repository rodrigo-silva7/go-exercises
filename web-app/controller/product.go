package controller

import (
	"html/template"
	"net/http"
   "web-app/model"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
   templates.ExecuteTemplate(w, "Index", model.GetProducts())
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
   templates.ExecuteTemplate(w, "New", nil)
}
