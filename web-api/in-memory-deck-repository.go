package main

import (
	"errors"
	"fmt"
)

type repo struct {
	decks []*Deck
}

func NewInMemoryDeckRepository() DeckRepository {
	return &repo{}
}

func (r *repo) Save(deck *Deck) (*Deck, error) {
	r.decks = append(r.decks, deck)
	fmt.Println("Saved deck with ID ", deck.Id)
	return deck, nil
}

func (r *repo) Get(id string) (*Deck, error) {
	return r.getDeck(id)
}

func (r *repo) DrawCards(id string, amount int) ([]card, error) {
	fmt.Println("Repo drawCards")
	deck, err := r.getDeck(id)

	if err != nil {
		return nil, DeckNotFoundError("Deck not found")
	}

	return deck.Draw(amount)
}

func (r *repo) getDeck(id string) (*Deck, error) {
	for _, deck := range r.decks {
		if deck.Id == id {
			return deck, nil
		}
	}

	return nil, errors.New("Deck not found")
}
