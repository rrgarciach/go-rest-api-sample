package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(addr, port string, db int) {
	fmt.Println("Initialized!", db)

	a.Router = mux.NewRouter()
	a.initializeRoutes()

	redisClient := Redis{}
	redisClient.Initialize(addr, port, db)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/xsltproc", a.hello).Methods("GET")
}

func (a *App) hello(w http.ResponseWriter, r *http.Request) {
	xsltProc := XsltProc{}
	response := xsltProc.transform()
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, response []byte) {
	// response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(code)
	w.Write(response)
}
