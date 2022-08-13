package router

import (
   "net/http"
   "web-app/controller"
)

func LoadRoutes() {
   http.HandleFunc("/", controller.Index)

   http.HandleFunc("/product", controller.NewProduct)
   http.HandleFunc("/insert", controller.Insert)
   http.HandleFunc("/update", controller.Update)
   http.HandleFunc("/delete", controller.Delete)
   http.HandleFunc("/edit", controller.Edit)
}

