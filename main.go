package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test simple HTTP server \n"))
}
func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP server on port 777")
	log.Fatal(http.ListenAndServe(":777", nil))
}
