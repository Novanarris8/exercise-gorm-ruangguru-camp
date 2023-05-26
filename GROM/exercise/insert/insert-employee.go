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
	jumlah := [5]Employee{{Name: "Aditira", Address: "test@gmail.com", Age: 23, Birthday: "1998-02-21", level: "Staff", id_departement: 1},
		{Name: "Aditira1", Address: "test@gmail.com", Age: 24, Birthday: "1998-02-22", level: "Supervisor", id_departement: 2},
		{Name: "Aditira2", Address: "test@gmail.com", Age: 25, Birthday: "1998-02-23", level: "Manager", id_departement: 3},
		{Name: "Aditira3", Address: "test@gmail.com", Age: 21, Birthday: "1998-02-24", level: "Staff", id_departement: 4},
		{Name: "Aditira4", Address: "test@gmail.com", Age: 22, Birthday: "1998-02-25", level: "Staff", id_departement: 5}}
	hasil := db.Create(&jumlah)

	fmt.Println("Rows: ", hasil.RowsAffected)
	for _, Employee := range jumlah {
		fmt.Println(Employee.ID, Employee.Name, Employee.Address, Employee.Age, Employee.Birthday, Employee.level, Employee.id_departement)
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Insert Success")
}
