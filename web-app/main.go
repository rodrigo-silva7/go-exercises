package main

import (
	"fmt"
	"net/http"
	"web-app/router"
)

var (
   PORT = 8000
)

func main() {
   router.LoadRoutes()
   fmt.Println("Servidor será iniciado na porta",PORT);
   http.ListenAndServe(fmt.Sprint(":",PORT), nil)   
}

