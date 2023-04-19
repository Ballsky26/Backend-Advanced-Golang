package main

import (
	"database/sql"
	_ "database/sql" // add this
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // add this
)

type newStudent struct {
	Student_id       uint64 `json:"student_id" binding:"required"`
	Student_name     string `json:"student_name" binding:"required"`
	Student_age      uint64 `json:"student_age" binding:"required"`
	Student_address  string `json:"student_address" binding:"required"`
	Student_phone_no string `json:"student_phone_no" binding:"required"`
}

func setupRouter() *gin.Engine {
	conn := "postgresql://postgres:1234@127.0.0.1/postgres?sslmode=disable"
	_, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
