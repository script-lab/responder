package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang")
}

func main() {
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
