package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sprint1_finalTask/internal/api/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"text": "hello!"})
	})

	router.POST("/api/v1/calculate", handlers.CalcMiddleware(), handlers.CalcHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
