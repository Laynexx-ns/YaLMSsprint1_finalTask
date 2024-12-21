package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"sprint1_finalTask/internal/api/services"
	"strings"
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

		//у меня калькулятор не обрабатывает / 0, но да ладно
		if strings.Contains(req.Expression, "/0") || strings.Contains(req.Expression, "/ 0") {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Dividing by zero"})
			return
		}

		//брекеты
		if !BracketsValidation(req.Expression) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Incorrect brackets"})
			return
		}

		//проверка на содержание мусора
		valid, err := regexp.MatchString("^[0-9)(*/+-]+$", req.Expression)
		if err != nil || !valid {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid characters in expression!!!"})
			return
		}

		c.Set("expression", req.Expression)

		c.Next()
	}
}

func CalcHandler(c *gin.Context) {
	log.Printf("Handler working")

	expression, _ := c.Get("expression")

	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic recovered")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}()

	ans, err := services.Calculate(expression.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка сервера"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": ans})
}

func BracketsValidation(exp string) bool {
	leftcount := 0
	rightcount := 0

	for _, r := range exp {
		if r == '(' {
			leftcount++
		}
		if r == ')' {
			rightcount++
		}
	}

	if leftcount != rightcount {
		return false
	}
	return true
}
