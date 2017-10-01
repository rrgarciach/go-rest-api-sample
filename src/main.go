package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var PORT, REDIS_HOST, REDIS_PORT, REDIS_DB string

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3030"
	}
	if REDIS_HOST = os.Getenv("REDIS_HOST"); REDIS_HOST == "" {
		REDIS_HOST = "localhost"
	}
	if REDIS_PORT = os.Getenv("REDIS_PORT"); REDIS_PORT == "" {
		REDIS_PORT = "6379"
	}
	if REDIS_DB = os.Getenv("REDIS_DB"); REDIS_DB == "" {
		REDIS_DB = "0"
	}

	host := REDIS_HOST
	port := REDIS_PORT
	db, err := strconv.ParseInt(REDIS_DB, 10, 0)
	if err != nil {
		fmt.Println(db, err)
	}

	a := App{}
	a.Initialize(host, port, int(db))

	a.Run(":" + PORT)
}
