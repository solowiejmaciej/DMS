package services

import (
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func GenerateCodeForUser(userId uint, client *redis.Client) (string, error) {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	key := fmt.Sprintf("%d-2facode-%s", userId, time.Now().Format("2006-01-02"))
	err := client.Set(key, code, 15*time.Minute).Err()
	if err != nil {
		log.Error("Error while setting code to redis", err)
		return "", err
	}
	log.Info("Code generated for user ", userId, " is ", code)

	return code, nil
}

func CheckCodeForUser(userId uint, code string, client *redis.Client) bool {
	key := fmt.Sprintf("%d-2facode-%s", userId, time.Now().Format("2006-01-02"))
	val, err := client.Get(key).Result()
	if err != nil {
		log.Error("Error while getting code from redis", err)
		return false
	}

	if val == code {
		return true
	}

	return false
}
