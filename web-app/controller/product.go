package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web-app/model"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
   templates.ExecuteTemplate(w, "Index", model.GetProducts())
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
   templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
   if r.Method == "POST" {
      nome := r.FormValue("nome")
      descricao := r.FormValue("descricao")
      quantidade := r.FormValue("quantidade")
      preco := r.FormValue("preco")

      precoFloat, err := strconv.ParseFloat(preco, 64) 
      if err != nil {
         log.Println("ERROR Não foi possível converter preço. error [%s]", err)
      }

      quantidadeInt, err := strconv.Atoi(quantidade) 
      if err != nil {
         log.Println("ERROR Não foi possível converter quantidade. error [%s]", err)
      }

      model.CreateProduct(nome, descricao, precoFloat, quantidadeInt)
   }

   http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
   if r.Method == "POST" {
      id := r.FormValue("id")
      nome := r.FormValue("nome")
      descricao := r.FormValue("descricao")
      quantidade := r.FormValue("quantidade")
      preco := r.FormValue("preco")

      precoFloat, err := strconv.ParseFloat(preco, 64) 
      if err != nil {
         log.Println("ERROR Não foi possível converter preço. error [%s]", err)
      }

      quantidadeInt, err := strconv.Atoi(quantidade) 
      if err != nil {
         log.Println("ERROR Não foi possível converter quantidade. error [%s]", err)
      }

      idInt, err := strconv.Atoi(id) 
      if err != nil {
         log.Println("ERROR Não foi possível converter id. error [%s]", err)
      }

      model.UpdateProduct(idInt, quantidadeInt, nome, descricao, precoFloat)

   }
   http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
   productId := r.URL.Query().Get("id")
   model.DeleteProduct(productId)

   http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
   product := model.EditProduct(r.URL.Query().Get("id"))

   templates.ExecuteTemplate(w, "Edit", product)
}
