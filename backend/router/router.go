package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	router.GET("/health", healthCheck)

	router.Run(":8080")
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "API is healthy!",
	})
}
