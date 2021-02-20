package main

import (
	"crypto/tls"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rncn/heeeserver/handler"
)

var (
	address         = flag.String("address", ":8080", "Server Listen Address")
	tsladdress      = flag.String("tsladdress", ":8081", "Server Listen Address(TSL)")
	certpemfilepath = flag.String("certfile", "localhost.pem", "Server Listen Address(TSL)")
	keyfilepath     = flag.String("keyfile", "localhost-key.pem", "Server Listen Address(TSL)")
)

//HttpServer is Redirect to HTTPS
func LaunchHTTPServer() {
	http.ListenAndServe(*address, nil)
}

//HttpsServer is Main Server
func LaunchHTTPSServer(router *mux.Router) int {
	//Load Certfile and Keyfile
	certPem, err := ioutil.ReadFile(*certpemfilepath)
	keyPem, err := ioutil.ReadFile(*keyfilepath)

	//CertConfig
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	tlsServe := &http.Server{
		Handler:      router,
		Addr:         *tsladdress,
		TLSConfig:    cfg,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	//ListenAndServe
	if err := tlsServe.ListenAndServeTLS("", ""); err != nil {
		panic(err)
	}

	return 0
}

//Router is Router
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.HelloWorldHandler)

	return router
}

func main() {
	flag.Parse()
	//Router
	router := Router()

	//Launching HTTP Server(Redirecting HTTPS server)
	if err := LaunchHTTPSServer(router); err != 0 {
		panic(err)
	}
}
