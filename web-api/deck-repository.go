package main

type DeckRepository interface {
	Save(deck *Deck) (*Deck, error)
}
