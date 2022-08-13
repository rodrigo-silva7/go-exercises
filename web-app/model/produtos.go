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

   productsQuery, err := connection.Query("select * from produtos order by id asc")
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

func CreateProduct(nome, descricao string, preco float64, quantidade int) {
   connection := db.OpenDatabaseConnection()

   insert, err := connection.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
   if err != nil {
      panic(err.Error())
   }

   insert.Exec(nome, descricao, preco, quantidade)
   defer connection.Close()
}

func DeleteProduct(id string) {
   connection := db.OpenDatabaseConnection()

   deleteQuery, err := connection.Prepare("delete from produtos where id = $1")
   if err != nil {
      panic(err.Error())
   }

   deleteQuery.Exec(id)
   defer connection.Close()
}

func EditProduct(id string) Product {
   connection := db.OpenDatabaseConnection()

   findProduct, err := connection.Query("Select * from produtos where id = $1", id)
   if err != nil {
      panic(err.Error())
   }
   product := Product{}
   for findProduct.Next() {
      var id, quantidade int
      var nome, descricao string
      var preco float64

      err = findProduct.Scan(&id, &nome, &descricao, &preco, &quantidade) 
      if err != nil {
         panic(err.Error())
      }

      product.Nome = nome
      product.Id = id
      product.Quantidade = quantidade
      product.Preco = preco
      product.Descricao = descricao
   }
   defer connection.Close()
   return product
}

func UpdateProduct(id, quantidade int, nome, descricao string, preco float64) {
   connection := db.OpenDatabaseConnection()

   updateQuery, err := connection.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
   if err != nil {
      panic(err.Error())
   }

   updateQuery.Exec(nome, descricao, preco, quantidade, id)
   defer connection.Close()


}
