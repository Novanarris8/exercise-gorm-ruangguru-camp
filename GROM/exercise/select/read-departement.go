package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Departement struct {
	id               uint
	nama_departement string
}

func main() {
	dsn := "host=localhost user=postgres password=novan8 dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var departement Departement
	db.First(&departement, 1)
	fmt.Println(departement)

}
