package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"sprint1_finalTask/internal/api/services"
)

type CalculateRequest struct {
	Expression string `json:"expression" binding:"required"`
}

func CalcMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

		var req CalculateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}
		valid, err := regexp.MatchString("^[0-9)(*/+-]+$", req.Expression)
		if err != nil || !valid {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid characters in expression"})
			return
		}

		c.Set("expression", req.Expression)

		c.Next()
	}
}

func CalcHandler(c *gin.Context) {
	log.Printf("Handler working")

	expression, _ := c.Get("expression")

	ans, err := services.Calculate(expression.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": ans})
}
