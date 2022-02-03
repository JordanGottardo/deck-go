package main

import "fmt"

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
