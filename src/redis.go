package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	RedisClient *redis.Client
}

func (redisClient *Redis) Initialize(addr, port string, db int) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: "",
		DB:       0,
	})
	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)
	redisClient.ExampleClient(RedisClient)
}

func (redisClient *Redis) ExampleClient(RedisClient *redis.Client) {
	err := RedisClient.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisClient.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := RedisClient.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}
