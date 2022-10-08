package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CreateConnection() (*redis.Client, context.Context) {
	ctx := context.Background()
	uri := viper.GetString("cache.redis.uri")
	opt, err := redis.ParseURL(uri)
	if err != nil {
		logrus.Error("Error parsing redis url ", uri)
		panic(err)
	}

	rdb := redis.NewClient(opt)

	return rdb, ctx
}
