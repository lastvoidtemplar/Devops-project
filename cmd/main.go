package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const HelloWorld = "Hello World"
const Port = 80

var Hostname string = ""

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("404 Not Found: %s", r.URL.Path)))
		return
	}

	log.Println("Hello World path was hit!")
	w.Write([]byte(HelloWorld))
}

func HostnameHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hostname path was hit!")
	w.Write([]byte(Hostname))
}

func main() {
	Hostname = os.Getenv("HOSTNAME")

	if Hostname == "" {
		log.Fatal("Missing ENV HOSTNAME")
	}

	log.Printf("Hostname: %s\n", Hostname)
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloWorldHandler)
	mux.HandleFunc("/hostname", HostnameHandler)

	log.Printf("Server is listening on port %d\n", Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), mux); err != nil {
		log.Fatalf("Error occured when server was listening on post %d: %s\n", Port, err.Error())
	}
}
