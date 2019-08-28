package bmredis

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"sync"
)

var onceConfig sync.Once
var redisClient *redis.Client

func GetRedisClient() *redis.Client {

	onceConfig.Do(func() {
		host := os.Getenv("BM_REDIS_HOST")
		port := os.Getenv("BM_REDIS_PORT")
		password := os.Getenv("BM_REDIS_PASS")
		dbStr := os.Getenv("BM_REDIS_DB")

		db, err := strconv.Atoi(dbStr)
		bmerror.PanicError(err)

		addr := host + ":" + port

		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // no password set
			DB:       db,  // use default DB
		})
		redisClient = client
	})

	return redisClient
}
