package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type App struct {
	Router      *mux.Router
	RedisClient *redis.Client
}

func (a *App) Initialize(addr string, port, db int) {
	print("Initialize!")
	a.Router = mux.NewRouter()
	a.initializeRoutes()
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
