package main

import (
	_ "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.New()

	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "true",
		})
	})

	port := "4000"
	server.Run(":" + port)
}
