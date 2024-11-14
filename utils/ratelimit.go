package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mu-wahba/url-shortener/databases"

	"github.com/gin-gonic/gin"
)

// var err error

func ShouldPass(c *gin.Context, ip string) string {
	//define vars
	api_quota := os.Getenv("API_QUOTA")
	api_ttl, _ := strconv.Atoi(os.Getenv("API_LIMIT_PERIOD"))

	//connect to db
	rdb, err := databases.ConnectDB(3)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"msg": "couldn't connect to redis"})
		return "notok"
	}
	//check if the ip exists
	if !CheckInRedis(3, ip) {
		fmt.Println("notexist")
		//if not exist, create it and set TTL ip:API_QUOTA:TTL
		err = rdb.Set(databases.Ctx, ip, api_quota, time.Duration(api_ttl)*60*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"msg": "couldn't set in db "})
			return "notok"
		}

	} else { //if exist, get quota and ttl , if possible decrease the quota by one and update in db
		// fmt.Println("exit")

		user_quota, err := rdb.Get(databases.Ctx, ip).Result()
		fmt.Println("user_quota", user_quota)
		user_quota_int, _ := strconv.Atoi(user_quota)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"msg": "couldn't Get from db "})
			return "notok"
		}
		user_ttl := rdb.TTL(databases.Ctx, ip).Val()
		//decrease by one and set in redis again
		user_quota_int = user_quota_int - 1

		if user_quota_int <= 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg":               "Rate limit excceded ",
				"remaining_minutes": user_ttl / time.Nanosecond / time.Minute,
			})
			return "notok"
		}
		err = rdb.Set(databases.Ctx, ip, user_quota_int, user_ttl).Err()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"msg": "couldn't set in db "})
			return "notok"
		}
	}

	return "ok"
}
