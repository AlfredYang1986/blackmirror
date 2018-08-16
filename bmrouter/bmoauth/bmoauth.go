package bmoauth

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func CheckToken(token string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	_, err := client.Get(token).Result()
	if err == redis.Nil {
		return errors.New("token not exist")
	} else if err != nil {
		panic(err)
	} else {
		return nil
	}
}

func PushToken(token string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	pipe := client.Pipeline()

	incr := pipe.Incr(token)
	pipe.Expire(token, 7*24*time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
	return err
}
