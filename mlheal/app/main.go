package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	httplogger "github.com/osunac/go-httplogger"
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

func pastDates(start, stop time.Time, step time.Duration) chan time.Time {
	c := make(chan time.Time)
	go func() {
		for t := start; t.Before(stop); t = t.Add(step) {
			c <- t
		}
		close(c)
	}()

	return c
}

func main() {
	stop := time.Now()
	start := stop.Add(-3 * 24 * time.Hour)
	step := 24 * time.Hour
	duration := 100000

	for t := range pastDates(start, stop, step) {
		log.Printf("HTTP - %s - - %s \"%s %s %s\" %d %d \"%s\" \"%s\" %dus\n",
			"127.0.0.1",
			t.Format("[02/Jan/2006:15:04:05 -0700]"),
			//"[02/Jan/2006:15:04:05 -0700]",
			"GET",
			"/",
			"HTTP/1.1",
			200,
			13,
			"",
			"App Client",
			//time.Since(t)/1000,
			duration,
		)
	}

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", hello)
	serveMux.HandleFunc("/break", breakApp)
	srv := http.Server{
		Addr:    ":8080",
		Handler: httplogger.HTTPLogger(serveMux),
	}

	log.Fatal(srv.ListenAndServe())
}
