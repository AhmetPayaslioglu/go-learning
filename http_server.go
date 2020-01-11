package main

import (
	"fmt"
	"net/http"
	"log"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You requested: " + request.URL.Path)
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("Error creating server. ", err)
	}
}
