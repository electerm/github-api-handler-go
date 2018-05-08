package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "os"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	//main
	w.Write([]byte("ok"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Print("ddd")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
