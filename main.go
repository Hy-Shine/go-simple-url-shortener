package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hy-shine/go-url-shortener/handler"
	"github.com/hy-shine/go-url-shortener/store"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey, Go URL shortener!",
		})
	})

	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortURL(c)
	})

	router.GET("/:shortURL", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := router.Run(":9100")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
