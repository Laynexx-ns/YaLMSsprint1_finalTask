package application

import (
	"github.com/gin-gonic/gin"
	"log"
	"sprint1_finalTask/internal/api/handlers"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"text": "hello!"})
	})

	router.POST("/api/v1/calculate", handlers.CalcMiddleware(), handlers.CalcHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
