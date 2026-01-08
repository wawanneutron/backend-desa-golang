package main

import (
	"backend-go/backend-desa/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config dari .env
	config.LoadEnv()

	// inisialisasi gin
	router := gin.Default()

	// membuat router dengan method GET
	router.GET("/", func(c *gin.Context) {
		// return dengan response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
			"data": []interface{}{
				gin.H{"name": "Alice", "age": 30},
				gin.H{"name": "Bob", "age": 25},
				gin.H{"name": "Charlie", "age": 35},
				gin.H{"name": "Jokowi", "age": 52},
			},
		})
	})

	// mulai server dengan port 3000
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
