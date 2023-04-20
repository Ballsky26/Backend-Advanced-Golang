package main

import (
	"database/sql"
	_ "database/sql" // add this
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // add this
)

type newStudent struct {
	Student_id       uint64 `json:"student_id" binding:"required"`
	Student_name     string `json:"student_name" binding:"required"`
	Student_age      uint64 `json:"student_age" binding:"required"`
	Student_address  string `json:"student_address" binding:"required"`
	Student_phone_no string `json:"student_phone_no" binding:"required"`
}

func rowToStruct(rows *sql.Rows, dest interface{}) error {
	destv := reflect.ValueOf(dest).Elem()

	args := make([]interface{}, destv.Type().Elem().NumField())

	for rows.Next() {
		rowp := reflect.New(destv.Type().Elem())
		rowv := rowp.Elem()

		for i := 0; i < rowv.NumField(); i++ {
			args[i] = rowv.Field(i).Addr().Interface()
		}

		if err := rows.Scan(args...); err != nil {
			return err
		}

		destv.Set(reflect.Append(destv, rowv))
	}

	return nil
}

func postHandler(c *gin.Context, db *sql.DB) {
	var newStudent newStudent

	if c.Bind(&newStudent) == nil {
		_, err := db.Exec("insert into students values ($1,$2,$3,$4,$5)", newStudent.Student_id, newStudent.Student_name, newStudent.Student_age, newStudent.Student_address, newStudent.Student_phone_no)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "success create"})
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
}

func getAllHandler(c *gin.Context, db *sql.DB) {
	var newStudent []newStudent

	row, err := db.Query("select * from students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	rowToStruct(row, &newStudent)

	if newStudent == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newStudent})

}

func getHandler(c *gin.Context, db *sql.DB) {
	var newStudent []newStudent

	studentId := c.Param("student_id")

	row, err := db.Query("select * from students where student_id = $1", studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	rowToStruct(row, &newStudent)

	if newStudent == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newStudent})
}

func putHandler(c *gin.Context, db *sql.DB) {
	var newStudent newStudent

	studentId := c.Param("student_id")

	if c.Bind(&newStudent) == nil {
		_, err := db.Exec("update students set student_name=$1 where student_id=$2", newStudent.Student_name, studentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "success update"})
	}

}

func delHandler(c *gin.Context, db *sql.DB) {
	studentId := c.Param("student_id")

	_, err := db.Exec("delete from students where student_id=$1", studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success delete"})

}

func setupRouter() *gin.Engine {
	conn := "postgresql://postgres:1234@127.0.0.1/postgres?sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	r := gin.Default()

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"status": "success",
	// 		"value":  "hello world",
	// 	})
	// })

	// r.POST("/student", func(ctx *gin.Context) {
	// 	postHandler(ctx, db)
	// })

	// r.GET("/student", func(ctx *gin.Context) {
	// 	getAllHandler(ctx, db)
	// })

	// r.GET("/student/:student_id", func(ctx *gin.Context) {
	// 	getHandler(ctx, db)
	// })

	// r.PUT("/student/:student_id", func(ctx *gin.Context) {
	// 	putHandler(ctx, db)
	// })

	// r.DELETE("/student/:student_id", func(ctx *gin.Context) {
	// 	delHandler(ctx, db)
	// })

	return r

}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&newStudent{})

	data := newStudent{}
	if db.Find(&data).RecordNotFound() {
		fmt.Println("=================== Run Seeder User ======================")
		seederUser(db)
	}
}

func seederUser(db *gorm.DB) {
	data := newStudent{
		Student_id:       1,
		Student_name:     "Dono",
		Student_age:      20,
		Student_address:  "Jakarta",
		Student_phone_no: "0123456789",
	}

	db.Create(&data)
}

func main() {
	r := setupRouter()

	r.Run()

}
