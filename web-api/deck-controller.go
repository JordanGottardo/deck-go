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
	DrawCards(resp http.ResponseWriter, req *http.Request)
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

func (c *controller) DrawCards(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	deckId := mux.Vars(req)["id"]
	cardsToDraw, err := strconv.Atoi(req.URL.Query().Get("amount"))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode((ServiceError{Message: "Cannot parse amount"}))
		fmt.Println(err)
		return
	}

	drawnCards, err := deckService.DrawCards(deckId, cardsToDraw)
	if err != nil {
		fmt.Println("Error is not nil")
		switch err.(type) {
		case DeckNotFoundError:
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode((ServiceError{Message: "Deck not found"}))
			return
		case NotEnoughCardsError:
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode((ServiceError{Message: "Invalid amount: not enough cards to draw"}))
			return
		}
	}

	result, err := json.Marshal(ToCardDtoSlice(drawnCards))
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

	getCardsDto := make([]CardDto, 0)
	for _, card := range deck.cards {
		getCardsDto = append(getCardsDto, ToCardDto(card))
	}
	getDeckDto.Cards = getCardsDto

	return getDeckDto
}

func ToCardDtoSlice(cards []card) []CardDto {
	cardDtoSlice := make([]CardDto, 0)

	for _, card := range cards {
		cardDtoSlice = append(cardDtoSlice, ToCardDto(card))
	}

	return cardDtoSlice
}

func ToCardDto(card card) CardDto {
	return CardDto{
		Value: card.value,
		Suit:  card.suit,
		Code:  string(card.value[0]) + string(card.suit[0]),
	}
}

type GetDeckDto struct {
	Id         string    `json:"deck_id"`
	Shuffled   bool      `json:"shuffled"`
	Remanining int       `json:"remaining"`
	Cards      []CardDto `json:"cards"`
}

type CardDto struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}
