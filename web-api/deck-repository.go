package main

import "fmt"

type deckRepository interface {
	save(deck *deck) *deck
}

type repo struct {
	decks []deck
}

func newInMemoryDeckRepository() deckRepository {
	return &repo{}
}

func (r *repo) save(deck *deck) *deck {
	r.decks = append(r.decks, *deck)
	fmt.Println("Saved deck with ID ", deck.Id)
	fmt.Println("Now decks are ", r.decks)
	return deck
}
