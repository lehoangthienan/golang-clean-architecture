package redis

import (
	"log"

	redis "github.com/go-redis/redis/v7"
)

func NewRedisClient(addr string) (*redis.Client, func()) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	return client, func() {
		err = client.Close()
		if err != nil {
			log.Println("Failed to close redis connection by error", err)
		}
		log.Println("Close redis connection")
	}
}
