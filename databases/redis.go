package databases

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

// var err error

func ConnectDB(dbnumber int) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       dbnumber,
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		// Return nil client and the error
		return nil, err
	}

	return rdb, nil

}
