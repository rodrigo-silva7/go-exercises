package model

import (
   "web-app/db"
)

type Product struct {
   Id int
   Nome, Descricao string
   Preco float64
   Quantidade int
}

func GetProducts() []Product {
   connection := db.OpenDatabaseConnection()

   productsQuery, err := connection.Query("select * from produtos")
   if err != nil {
      panic(err.Error())
   }
   
   p := Product{}
   products := []Product{}

   for productsQuery.Next(){
      var id, quantidade int
      var nome, descricao string
      var preco float64

      err = productsQuery.Scan(&id, &nome, &descricao, &preco, &quantidade)
      if err != nil {
         panic(err.Error())
      }

      p.Nome = nome
      p.Id = id
      p.Descricao = descricao
      p.Preco = preco
      p.Quantidade = quantidade

      products = append(products, p)
   }
   defer connection.Close()
   return products
}

