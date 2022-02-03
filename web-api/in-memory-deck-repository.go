package main

import (
	"errors"
	"fmt"
)

type repo struct {
	decks []Deck
}

func NewInMemoryDeckRepository() DeckRepository {
	return &repo{}
}

func (r *repo) Save(deck *Deck) (*Deck, error) {
	r.decks = append(r.decks, *deck)
	fmt.Println("Saved deck with ID ", deck.Id)
	return deck, nil
}

func (r *repo) Get(id string) (*Deck, error) {
	for _, deck := range r.decks{
		if deck.Id == id {
			return &deck, nil
		}
	}

	return nil, errors.New("Deck not found")
}
