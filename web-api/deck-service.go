package main

import (
	"errors"

	"github.com/google/uuid"
)

type DeckService interface {
	Validate(deck *Deck) error
	Create(createDeckDto CreateDeckDto) (*Deck, error)
	Get(id string) (*Deck, error)
}

type service struct{}

type CreateDeckDto struct {
	Shuffled      bool
	RequiredCards []RequiredCard
}

type RequiredCard = string

var (
	deckRepo DeckRepository
)

func NewDeckService(repository DeckRepository) DeckService {
	deckRepo = repository
	return &service{}
}

func (*service) Validate(deck *Deck) error {
	if deck == nil {
		return errors.New("Deck is nil")
	}
	return nil
}

func (*service) Create(createDeckDto CreateDeckDto) (*Deck, error) {
	deck := newDeck()
	if createDeckDto.Shuffled {
		deck.Shuffle()
	}
	deck.Id = uuid.NewString()
	return deckRepo.Save(&deck)
}

func (*service) Get(id string) (*Deck, error) {
	return deckRepo.Get(id)
}
