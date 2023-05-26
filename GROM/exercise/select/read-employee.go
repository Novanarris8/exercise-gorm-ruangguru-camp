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
	jumlah := [5]Departement{{id: 1, nama_departement: "departement1"},
		{id: 2, nama_departement: "departement2"},
		{id: 3, nama_departement: "departement3"},
		{id: 4, nama_departement: "departement4"},
		{id: 5, nama_departement: "departement5"}}
	hasil := db.Create(&jumlah)

	fmt.Println("Rows: ", hasil.RowsAffected)
	for _, Departement := range jumlah {
		fmt.Println(Departement.id, Departement.nama_departement)
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Insert Success")
}
