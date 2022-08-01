package main

import (
	"net/http"
   "web-app/router"
)

func main() {
   router.LoadRoutes()
   http.ListenAndServe(":8000", nil)   
}

