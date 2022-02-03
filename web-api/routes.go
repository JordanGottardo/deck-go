package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

var (
	deckRepo deckRepository = newInMemoryDeckRepository()
)

func createNewDeck(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	deck := newDeck(uuid.NewString())
	deckRepo.save(&deck)
	result, err := json.Marshal(deck)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling decks"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
