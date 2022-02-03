package main

import (
	"errors"

	"github.com/google/uuid"
)

type DeckService interface {
	Validate(deck *Deck) error
	Create(deck *Deck) (*Deck, error)
}

type service struct{}

var (
	r DeckRepository = NewInMemoryDeckRepository()
)

func NewDeckService() DeckService {
	return &service{}
}

func (*service) Validate(deck *Deck) error {
	if deck == nil {
		return errors.New("Deck is nil")
	}
	return nil
}

func (*service) Create(deck *Deck) (*Deck, error) {
	deck.Id = uuid.NewString()
	return r.Save(deck)
}
