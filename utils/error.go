package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, code int, message string, err error) {
	// 1. Log it once here
	if err != nil {
		fmt.Printf("[MACHINE LOG] %s: %v\n", message, err)
	} else {
		fmt.Printf("[MACHINE LOG] %s\n", message)
	}

	// 2. Send the JSON response
	c.JSON(code, gin.H{
		"error":   true,
		"message": message,
	})
}
