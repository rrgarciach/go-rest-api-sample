package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
	pubsub *redis.PubSub
}

func (redisClient *Redis) Initialize(addr, port string, db int) {
	fmt.Println("Initializing redis...")
	client := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	fmt.Println("Subscribing to xsltproc channel...")
	pubsub := client.Subscribe("xsltproc")

	for i := 0; i < 1; i = 0 {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		// fmt.Println(msg.Channel, msg.Payload)

		message := strings.Split(msg.Payload, "#")
		params := "'" + message[0] + "'"
		data := message[1]

		fetchXmlService := FetchXmlService{}
		err = fetchXmlService.createXmlFileFromString(data)
		if err != nil {
			log.Fatal(err)
		}
		xsltProc := XsltProc{}
		result := xsltProc.transformFromString(params)

		err = client.Publish("fetchxml", string(result[:])).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println("Published to fetchxml channel")
	}
}

// func (redisClient *Redis) PubSubConn() {
// 	fmt.Println("Subscribing to xsltproc channel...")
// 	redisClient.pubsub = redisClient.client.Subscribe("xsltproc")
// 	// defer pubsub.Close()
//
// 	// subscr, err := pubsub.ReceiveTimeout(time.Second)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(subscr)
//
// 	// err = client.Publish("xsltproc", "hello").Err()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	for i := 0; i < 1; i = 0 {
// 		msg, err := redisClient.pubsub.ReceiveMessage()
// 		if err != nil {
// 			fmt.Println("error while receiving message")
// 			panic(err)
// 		}
// 		fmt.Println(msg.Channel, msg.Payload)
// 	}
// 	// msgCh := pubsub.Channel()
// 	// fmt.Println(msgCh)
// }

func (redisClient *Redis) ExampleClient(RedisClient *redis.Client) {
	// key2 does not exists
}
