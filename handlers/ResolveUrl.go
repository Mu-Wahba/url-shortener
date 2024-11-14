package handlers

import (
	"log"
	"net/http"

	databases "github.com/mu-wahba/url-shortener/databases"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func ResolveUrl(c *gin.Context) {
	url := c.Param("url")
	db, err := databases.ConnectDB(1)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Error couldn't connect to redis server")
		return
	}
	// Get the value from Redis
	value, err := db.Get(databases.Ctx, url).Result()
	if err == redis.Nil {
		//key doesn't exist
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Key not exists",
		})

	} else if err != nil {
		log.Printf("Could not get key: %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "error",
		})

	} else {
		//key exists
		//get the TTL
		ttl, err := db.TTL(databases.Ctx, url).Result()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"msg": "something went wrong ",
			})
		}

		// c.JSON(http.StatusOK, gin.H{
		// 	"url":       url,
		// 	"short_url": value,
		// 	"ttl":       ttl.Seconds() / 60,
		// })
		if ttl.Abs() > 0 {
			c.Redirect(301, value)
		}
	}

}
