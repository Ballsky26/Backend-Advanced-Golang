package main

import (
	"fmt"

	// package used to read the .env file
	_ "github.com/lib/pq" // postgres golang driver

	models "Gorm/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func main() {

	var err error
	// DB, err = gorm.Open("mysql", "root:@/egommerce?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("postgres", "postgresql://postgres:1234@127.0.0.1/postgres?sslmode=disable")
	if err != nil {
		panic("Gagal konek ke db")
	}
	defer DB.Close()

	fmt.Println("Sukses Konek ke Db!")

	Migrate()

}

func Migrate() {
	DB.AutoMigrate(&models.Student{})
}
