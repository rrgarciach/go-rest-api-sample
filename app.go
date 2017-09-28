package main

import (
	"encoding/json"
	"log"
	"net/http"
  "fmt"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type App struct {
	Router      *mux.Router
	RedisClient *redis.Client
}

func (a *App) ExampleClient(RedisClient *redis.Client) {
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

func (a *App) Initialize(addr, port string, db int) {
  fmt.Println("Initialized!", db)
	a.Router = mux.NewRouter()
	a.initializeRoutes()

  RedisClient := redis.NewClient(&redis.Options{
    Addr: addr + ":" + port,
    Password: "",
    DB: 0,
  })
  pong, err := RedisClient.Ping().Result()
  fmt.Println(pong, err)
  a.ExampleClient(RedisClient)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/hello", a.hello).Methods("GET")
}

func (a *App) hello(w http.ResponseWriter, r *http.Request) {
	response := "HI!!!"
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
