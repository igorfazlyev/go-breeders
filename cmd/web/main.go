package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":4000"

type Application struct{}

func main() {
	//var app = Application{}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "the server is up and running")
	})
	fmt.Println("Starting web application on port", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("it crashed", err)
	}
}
