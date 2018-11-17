package bmredis

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/go-redis/redis"
)

func GetRedisClient() *redis.Client {

	redisConfigPath := "resource/redisconfig.json"
	redisConfig := bmconfig.BMGetConfigMap(redisConfigPath)

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
	return client
}
