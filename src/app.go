package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(addr, port string, db int) {
	a.Router = mux.NewRouter()
	a.initializeRoutes()

	// redisClient := Redis{}
	// redisClient.Initialize(addr, port, db)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/xsltproc", a.xsltprocXML).Methods("GET").Headers("Accept", "application/xml")
}

func (a *App) xsltprocXML(w http.ResponseWriter, r *http.Request) {
	fetchXml := FetchXml{}
	err := fetchXml.getFetchXml()
	if err != nil {
		log.Fatal(err)
	}

	xsltProc := XsltProc{}
	response := xsltProc.transform()

	respondWithXML(w, http.StatusOK, response)
}

func respondWithXML(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithJSON(w http.ResponseWriter, code int, response []byte) {
	// response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
