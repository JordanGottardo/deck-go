package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	deckService DeckService
)

type DeckController interface {
	CreateNewDeck(resp http.ResponseWriter, req *http.Request)
	GetDeck(resp http.ResponseWriter, req *http.Request)
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

func (c *controller) GetDeck(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	id := mux.Vars(req)["id"]
	fmt.Println(id)

	deck, err := deckService.Get(id)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"error": "Deck not found"}`))
		return
	}

	getDeckDto := ToGetDeckDto(deck)
	result, err := json.Marshal(getDeckDto)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling decks"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func ToGetDeckDto(deck *Deck) GetDeckDto {
	getDeckDto := GetDeckDto{
		Id:         deck.Id,
		Shuffled:   deck.IsShuffled,
		Remanining: deck.remainingCards(),
	}

	getCardsDto := make([]GetCardDto, 0)
	for _, card := range deck.cards {
		getCardsDto = append(getCardsDto, GetCardDto{
			Value: card.value,
			Suit:  card.suit,
			Code:  string(card.value[0]) + string(card.suit[0]),
		})
	}
	getDeckDto.Cards = getCardsDto

	return getDeckDto
}

type GetDeckDto struct {
	Id         string       `json:"deck_id"`
	Shuffled   bool         `json:"shuffled"`
	Remanining int          `json:"remaining"`
	Cards      []GetCardDto `json:"cards"`
}

type GetCardDto struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}
