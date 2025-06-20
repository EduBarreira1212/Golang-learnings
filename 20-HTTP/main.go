package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Root route!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home route"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users route"))
}

func main() {
	http.HandleFunc("/", root)

	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
