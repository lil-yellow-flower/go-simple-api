package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/sort", sortCards)

	router.Run("localhost:8080")
}

func sortCards(c *gin.Context) {
	input := c.Query("cards")

	validated := validator(input)
	output := sorter(validated)

	c.IndentedJSON(http.StatusOK, output)
}

func validator(input string) []string {
	return strings.Split(input, ",")
}

func sorter(input []string) []string {
	return input
}
