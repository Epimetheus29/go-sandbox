package handlers

import (
	"net/http"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
)

// FlipCasing toggles lowercase to uppercase and vice versa for an input string
func FlipCasing(c *gin.Context) {
	// Grab the text from the URL query parameter (e.g., /api/flip?text=HelloGo)
	inputText := c.Query("text")
	if inputText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'text' query parameter"})
		return
	}

	// Convert string to a slice of runes to support multi-byte UTF-8 safely
	runes := []rune(inputText)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"original": inputText,
		"flipped":  string(runes),
	})
}

// Fibonacci generates a sequence up to 'n' elements
func Fibonacci(c *gin.Context) {
	// Grab string parameter and safely parse it to an integer
	countStr := c.DefaultQuery("count", "10") // defaults to 10 if not provided
	count, err := strconv.Atoi(countStr)

	if err != nil || count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'count' parameter. Must be a positive integer."})
		return
	}

	// Guard against massive CPU loops in our sandbox
	if count > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Count too high for a sandbox! Keep it under 100."})
		return
	}

	// Generate sequence
	sequence := make([]int, count)
	if count > 0 {
		sequence[0] = 0
	}
	if count > 1 {
		sequence[1] = 1
	}
	for i := 2; i < count; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	c.JSON(http.StatusOK, gin.H{
		"count":    count,
		"sequence": sequence,
	})
}
