package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	SetJsonContentHeader(resp)
	shuffledParam := req.URL.Query()["shuffled"]
	shuffled := false
	if len(shuffledParam) > 0 {
		var err error
		shuffled, err = strconv.ParseBool(shuffledParam[0])

		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			WriteError(resp, "Invalid shuffled value provided")
			return
		}
	}

	cardsParam := req.URL.Query()["cards"]
	var requestedCards []string
	if len(cardsParam) > 0 {
		requestedCards = strings.Split(cardsParam[0], ",")
	}

	createDeckDto := CreateDeckDto{
		Shuffled:       shuffled,
		RequestedCards: ToRequestedCards(requestedCards),
	}
	deck, err := deckService.Create(createDeckDto)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		WriteError(resp, "Invalid card provided")
		return
	}

	result, err := json.Marshal(deck)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		WriteError(resp, "Error marshalling deck")
		return
	}

	fmt.Println("Created new deck with id ", deck.Id)
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func (c *controller) GetDeck(resp http.ResponseWriter, req *http.Request) {
	SetJsonContentHeader(resp)
	id := mux.Vars(req)["id"]

	deck, err := deckService.Get(id)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		WriteError(resp, "Deck not found")
		return
	}

	getDeckDto := ToGetDeckDto(deck)
	result, err := json.Marshal(getDeckDto)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		WriteError(resp, "Error marshalling deck")
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func (c *controller) DrawCards(resp http.ResponseWriter, req *http.Request) {
	SetJsonContentHeader(resp)
	deckId := mux.Vars(req)["id"]
	cardsToDraw, err := strconv.Atoi(req.URL.Query().Get("amount"))

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		WriteError(resp, "Cannot parse amount")
		return
	}

	drawnCards, err := deckService.DrawCards(deckId, cardsToDraw)
	if err != nil {
		switch err.(type) {
		case DeckNotFoundError:
			resp.WriteHeader(http.StatusNotFound)
			WriteError(resp, "Deck not found")
			return
		case NotEnoughCardsError:
			resp.WriteHeader(http.StatusBadRequest)
			WriteError(resp, "Invalid amount: not enough cards to draw")
			return
		}
	}

	result, err := json.Marshal(ToCardDtoSlice(drawnCards))
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		WriteError(resp, "Error marshalling deck")
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

// Utility methods

func SetJsonContentHeader(resp http.ResponseWriter) {
	resp.Header().Set("Content-type", "application/json")
}

func WriteError(resp http.ResponseWriter, errorMessage string) {
	json.NewEncoder(resp).Encode((ServiceError{Message: errorMessage}))
}

func ToGetDeckDto(deck *Deck) GetDeckDto {
	getDeckDto := GetDeckDto{
		Id:         deck.Id,
		Shuffled:   deck.IsShuffled,
		Remanining: deck.RemainingCards(),
	}

	getCardsDto := make([]CardDto, 0)
	for _, card := range deck.Cards {
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

func ToRequestedCards(cards []string) []RequestedCard {
	requestedCards := make([]RequestedCard, 0)

	for _, card := range cards {
		requestedCards = append(requestedCards, RequestedCard(card))
	}

	return requestedCards
}
