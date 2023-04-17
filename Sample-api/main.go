package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// ping
	r.GET("/ping", func(ctx *gin.Context) {
		// ctx.String(http.StatusOK, "Hello World")

		// ctx.JSON(http.StatusOK, map[string]interface{}{
		// 	"status": "success",
		// 	"value":  "Hello world",
		// })

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  "Hello world",
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
