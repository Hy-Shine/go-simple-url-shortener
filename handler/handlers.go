package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hy-shine/go-url-shortener/shortener"
	"github.com/hy-shine/go-url-shortener/store"
)

type URLCreateRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	// UserID  string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	var creationRequest URLCreateRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	shortURL := shortener.GenerateShortLink(creationRequest.LongURL) // creationRequest.UserID)
	store.SaveURLMapping(shortURL, creationRequest.LongURL)          //creationRequest.UserID)
	host := "localhost:9100/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortURL,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortURL := c.Param("shortURL")
	initialURL := store.RetrieveInitiaURL(shortURL)
	c.Redirect(302, initialURL)
}
