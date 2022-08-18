package main

import (
	"alura-rest-api/database"
	"alura-rest-api/router"
	"fmt"
)

var (
   PORT = "8080"
)

func main() {

   database.OpenDatabaseConnection()

   fmt.Println("Servidor iniciado na porta", PORT)
   router.HandleProvider()
}

