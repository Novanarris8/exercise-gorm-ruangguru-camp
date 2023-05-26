package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID             uint
	Name           string
	Address        string
	Age            uint8
	Birthday       string
	level          string
	id_departement int
}

func main() {
	dsn := "host=localhost user=postgres password=novan8 dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var employee Employee
	db.First(&employee, 1)
	fmt.Println(employee)

}
