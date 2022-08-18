package database

import (
   "log"
   "gorm.io/gorm"
   "gorm.io/driver/postgres"
)

var (
   DB *gorm.DB  
   err error
)

func OpenDatabaseConnection() {
   connectionString := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable Timezone=America/Sao_Paulo"
   DB, err = gorm.Open(postgres.Open(connectionString))
   if err != nil {
      log.Panic("Erro ao conectar com o banco de dados!")
   }
}
