package main


import (
	"github.com/mu-wahba/url-shortener/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

}

func main() {
	router := gin.Default()

	router.POST("/api/v2/url", handlers.ShortenUrl)
	router.GET("/:url", handlers.ResolveUrl)
	router.Run()

}
