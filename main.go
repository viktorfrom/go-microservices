package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"time"

	details "github.com/viktorfrom/go-microservices/details"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking app health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}

	json.NewEncoder(w).Encode(response)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App is up and running")

}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}

	IP, err := details.GetIP()

	fmt.Println(hostname, IP)

	response := map[string]string{
		"hostname": hostname,
		"ip":       IP.String(),
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	//r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/details", detailsHandler)
	r.HandleFunc("/", rootHandler)

	http.Handle("/", r)

	log.Println("Server has started")

	log.Fatal(http.ListenAndServe(":3000", r))
}
