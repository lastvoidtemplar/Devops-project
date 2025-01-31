package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const HelloWorld = "Hello World"
const Port = 3000

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(HelloWorld))
}

func HostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()

	if err != nil {
		log.Printf("Error occured when getting hostname from the kernel: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn`t get hostname!"))
		return
	}

	w.Write([]byte(hostname))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloWorldHandler)
	mux.HandleFunc("/hostname", HostnameHandler)

	log.Printf("Server is listening on port %d", Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), mux); err != nil {
		log.Fatalf("Error occured when server was listening on post %d: %s", Port, err.Error())
	}
}
