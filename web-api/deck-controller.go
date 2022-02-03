package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	shuffled, _ := strconv.ParseBool((req.URL.Query()["shuffled"][0]))
	fmt.Println(shuffled)
	createDeckDto := CreateDeckDto{
		Shuffled:      shuffled,
		RequiredCards: []RequiredCard{},
	}
	deck, _ := deckService.Create(createDeckDto)
	result, err := json.Marshal(deck)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling decks"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
