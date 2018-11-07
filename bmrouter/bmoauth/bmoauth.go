package bmoauth

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func CheckToken(token string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.100.174:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
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
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.100.174:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	pipe := client.Pipeline()

	pipe.Incr(token)
	pipe.Expire(token, 7*24*time.Hour)

	_, err := pipe.Exec()
	fmt.Println(token)
	return err
}
