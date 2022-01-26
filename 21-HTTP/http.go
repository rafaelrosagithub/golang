package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Test..."))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loading users page..."))
}

func main() {

	// HTTP is a communicaton protocol, basis of web communication
	// Client (make request) - server (process and send response)
	// Request - Response

	// Routes
	// URI - resource identifier
	// Methods - GET, POST, PUT, DELET

	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
