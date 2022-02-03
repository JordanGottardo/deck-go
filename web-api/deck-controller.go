package main

import (
	"encoding/json"
	"net/http"
)

var (
	deckService DeckService
)

type DeckController interface {
	CreateNewDeck(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewDeckController(service DeckService) DeckController {
	deckService = service
	return &controller{}
}

func (c *controller) CreateNewDeck(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	newDeck := newDeck()
	deck, _ := deckService.Create(&newDeck)
	result, err := json.Marshal(deck)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling decks"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
