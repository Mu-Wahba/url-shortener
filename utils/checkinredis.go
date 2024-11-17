package utils

import (
	"log"

	"github.com/mu-wahba/url-shortener/databases"
)

func CheckInRedis(dbnumber int, key string) bool {
	rdb, err := databases.ConnectDB(dbnumber)
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()
	value, _ := rdb.Get(databases.Ctx, key).Result()
	return value != ""
}
