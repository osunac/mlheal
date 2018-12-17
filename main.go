package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	httplogger "github.com/gleicon/go-httplogger"
)

var isBroken = false

func hello(w http.ResponseWriter, r *http.Request) {
	sleepTime := 100 * time.Millisecond
	responseMessage := "Hello, world\n"

	if isBroken {
		sleepTime += 2 * time.Second
		responseMessage = "Hlleo, wrlod\n"
	}

	time.Sleep(sleepTime)
	fmt.Fprintf(w, responseMessage)
}

func breakApp(w http.ResponseWriter, r *http.Request) {
	isBroken = true

	fmt.Fprintf(w, "Application is broken now.\n")
}

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", hello)
	serveMux.HandleFunc("/break", breakApp)
	srv := http.Server{
		Addr:    ":8080",
		Handler: httplogger.HTTPLogger(serveMux),
	}

	log.Fatal(srv.ListenAndServe())
}
