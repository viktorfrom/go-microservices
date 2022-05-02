package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	bookapi "github.com/viktorfrom/go-microservices/api"
	details "github.com/viktorfrom/go-microservices/details"
)

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
	// Init Router
	r := mux.NewRouter()

	// Mock Data - $todo - implement DB
	bookapi.Books = append(bookapi.Books, bookapi.Book{ID: "1", ISBN: "448743", Title: "Book One",
		Author: &bookapi.Author{Firstname: "John", Lastname: "Doe"}})
	bookapi.Books = append(bookapi.Books, bookapi.Book{ID: "2", ISBN: "847564", Title: "Book Two",
		Author: &bookapi.Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", bookapi.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", bookapi.GetBook).Methods("GET")
	r.HandleFunc("/api/books", bookapi.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", bookapi.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", bookapi.DeleteBook).Methods("DELETE")

	r.HandleFunc("/details", detailsHandler)

	r.HandleFunc("/", rootHandler)
	http.Handle("/", r)

	log.Println("Server has started")

	log.Fatal(http.ListenAndServe(":3000", r))
}
