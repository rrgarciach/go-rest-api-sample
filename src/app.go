package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router      *mux.Router
	redisClient Redis
}

func (a *App) Initialize(addr, port string, db int) {
	a.Router = mux.NewRouter()
	// a.initializeRoutes()

	redisClient := Redis{}
	redisClient.Initialize(addr, port, db)
	// a.redisClient.PubSubConn()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/v1/xsltproc", a.xsltprocFetchXml).Methods("GET").Headers("Accept", "text/xml")
	a.Router.HandleFunc("/api/v1/xsltproc", a.xsltprocXml).Methods("POST").Headers("Accept", "text/xml", "Content-Type", "text/xml")
}

func (a *App) xsltprocFetchXml(w http.ResponseWriter, r *http.Request) {
	fetchXmlService := FetchXmlService{}

	fmt.Println("Request received:", r.Method, r.URL)

	err := fetchXmlService.getFetchXml()
	if err != nil {
		log.Fatal(err)
	}

	xsltProc := XsltProc{}
	params := r.URL.Query()
	response := xsltProc.transform(params)

	respondWithXML(w, http.StatusOK, response)
}

func (a *App) xsltprocXml(w http.ResponseWriter, r *http.Request) {
	fetchXmlService := FetchXmlService{}

	fmt.Println("Request received:", r.Method, r.URL)

	err := fetchXmlService.createXmlFile(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	xsltProc := XsltProc{}
	params := r.URL.Query()
	response := xsltProc.transform(params)

	respondWithXML(w, http.StatusOK, response)
}

func respondWithXML(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithJSON(w http.ResponseWriter, code int, response []byte) {
	// response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
