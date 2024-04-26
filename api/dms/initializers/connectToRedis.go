package initializers

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"os"
)

var RedisClient *redis.Client

func ConnectToRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Info("Connected to Redis")
}
