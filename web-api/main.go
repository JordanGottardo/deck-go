package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8000"
	router := mux.NewRouter()
	router.HandleFunc("/decks", createNewDeck).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}
