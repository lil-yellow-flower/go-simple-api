package main

import (
	"net/http"
	"strings"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	cacheHandler := cache.CacheByRequestURI(memoryStore, 2*time.Second)
	router.GET("/sort", cacheHandler, sortCards)

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
	var resultString []string

	dict := make(map[string]string)
	from := make(map[string]bool)
	to := make(map[string]bool)

	// separate
	for _, val := range input {
		pair := strings.Split(val, "-")
		f := pair[0]
		t := pair[1]

		from[f] = true
		to[t] = true

		dict[f] = t
	}

	// get starting point
	currentCard := struct {
		from string
		to   string
	}{
		from: "",
		to:   "",
	}

	for f, _ := range from {
		if to[f] == false {
			currentCard.from = f
			currentCard.to = dict[f]
			break
		}
	}
	resultString = append(resultString, currentCard.from+"-"+currentCard.to)

	// get rest and final destination
	for true {
		nextCardFrom := dict[currentCard.to]
		if nextCardFrom == "" {
			break
		}

		currentCard.from = currentCard.to
		currentCard.to = nextCardFrom

		resultString = append(resultString, currentCard.from+"-"+currentCard.to)
	}

	return resultString
}
