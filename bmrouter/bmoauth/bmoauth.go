package bmoauth

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmredis"
	"github.com/go-redis/redis"
	"time"
)

func CheckToken(token string) error {
	client := bmredis.GetRedisClient()
	defer client.Close()

	_, err := client.Get(token).Result()
	if err == redis.Nil {
		return errors.New("token not exist")
	} else if err != nil {
		//panic(err)
		fmt.Println(err.Error())
		return err
	} else {
		return nil
	}
}

func PushToken(token string) error {
	client := bmredis.GetRedisClient()
	defer client.Close()

	pipe := client.Pipeline()

	pipe.Incr(token)
	pipe.Expire(token, 365*24*time.Hour)

	_, err := pipe.Exec()
	fmt.Println(token)
	return err
}
