package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type Application struct{}

func main() {
	var app = Application{}
	srv := &http.Server{
		Addr:         port,
		Handler:      app.Routes(),
		IdleTimeout:  time.Second * 30,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	fmt.Println("starting web app on port", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
