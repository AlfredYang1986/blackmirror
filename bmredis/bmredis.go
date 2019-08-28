package bmredis

import (
	"blackmirror/bmconfighandle"
	"github.com/go-redis/redis"
	"os"
	"sync"
)

var onceConfig sync.Once
var redisClient *redis.Client

func GetRedisClient() *redis.Client {

	onceConfig.Do(func() {
		configPath := os.Getenv("BM_REDIS_CONF_HOME")
		redisConfig := bmconfig.BMGetConfigMap(configPath)

		host := redisConfig["Host"].(string)
		port := redisConfig["Port"].(string)
		addr := host + ":" + port
		password := redisConfig["Password"].(string)
		db := int(redisConfig["DB"].(float64))

		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // no password set
			DB:       db,  // use default DB
		})
		redisClient = client
	})

	return redisClient
}
