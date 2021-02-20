package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rncn/heeeserver/handler"
)

var (
	address = flag.String("address", ":8080", "Server Listen Address")
	tsladdress = flag.String("tsladdress", ":8081", "Server Listen Address(TSL)")
)

//Router is Router
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.HelloWorldHandler)

	return router
}

func main() {
	flag.Parse()

	router := Router()
	serve := &http.Server{
		Handler: router,
		Addr:    *address,
		//Server Access Timeout
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	tslServe := &http.Server{
		Handler: router,
		Addr: 
	}
	//ListenAndServe
	if err := serve.ListenAndServe(); err != nil {
		panic(err)
	}
}
