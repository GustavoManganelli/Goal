package router

import (
	_ "example.com/goal/backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Money organization API
// @version         1.0
// @description     A spents management for small groups service API in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @contact.email  goal.projectapp@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1
func Initialize() {
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", healthCheck)
	}

	r.Run(":8080")
}

// GetHealth         godoc
// @Summary Get api health
// @Description get api health status
// @Tags         health
// @Success      200
// @Router       /health [get]
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "API is healthy!",
	})
}
