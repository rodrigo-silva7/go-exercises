package router

import (
   "net/http"
   "web-app/controller"
)

func LoadRoutes() {
   http.HandleFunc("/", controller.Index)
}
