package main

import (
	"github.com/google/uuid"
)

type DeckService interface {
	Create(createDeckDto CreateDeckDto) (*Deck, error)
	Get(id string) (*Deck, error)
	DrawCards(deckId string, amount int) ([]card, error)
}

type service struct{}

type CreateDeckDto struct {
	Shuffled       bool
	RequestedCards []RequestedCard
}

type RequestedCard string

func (r *RequestedCard) GetValueAndSuite() (string, string) {
	requestedCard := string(*r)
	return string(requestedCard[0]), string(requestedCard[1])
}

var (
	deckRepo DeckRepository
)

func NewDeckService(repository DeckRepository) DeckService {
	deckRepo = repository
	return &service{}
}

func (*service) Create(createDeckDto CreateDeckDto) (*Deck, error) {
	var deck Deck
	if len(createDeckDto.RequestedCards) > 0 {
		cards, err := ToCards(createDeckDto.RequestedCards)

		if err != nil {
			return nil, err
		}

		deck = newDeckWithRequestedCards(cards)
	} else {
		deck = newDeck()
	}

	if createDeckDto.Shuffled {
		deck.Shuffle()
	}
	deck.Id = uuid.NewString()
	return deckRepo.Save(&deck)
}

func ToCards(requestedCards []RequestedCard) ([]card, error) {
	var cards []card

	for _, requestedCard := range requestedCards {
		value, suit := requestedCard.GetValueAndSuite()

		card, err := GetCard(value, suit)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}

func (*service) Get(id string) (*Deck, error) {
	return deckRepo.Get(id)
}

func (*service) DrawCards(deckId string, amount int) ([]card, error) {
	return deckRepo.DrawCards(deckId, amount)
}
