package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rncn/heeeserver/handler"
)

var (
	address    = flag.String("address", ":8080", "Server Listen Address")
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

	//CertConfig
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	tlsServe := &http.Server{
		Handler:   router,
		Addr:      *tsladdress,
		TLSConfig: cfg,
	}
	//ListenAndServe
	if err := tlsServe.ListenAndServe(); err != nil {
		panic(err)
	}
}
