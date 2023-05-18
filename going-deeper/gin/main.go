package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Animal struct {
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Breed string `json:"breed"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Animal{
			Name:  "Potter",
			Age:   13,
			Breed: "chandozo",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
