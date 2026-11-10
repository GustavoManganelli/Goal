package main

import "github.com/gin-gonic/gin"

func main() {
    // Cria uma instância do Gin no modo release
    router := gin.Default()

    // Endpoint GET /health
    router.GET("/health", healthCheck)

    // Inicia o servidor na porta 8080
    router.Run(":8080")
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "API is healthy!",
	})
}