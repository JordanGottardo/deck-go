package main

type DeckRepository interface {
	Save(deck *Deck) (*Deck, error)
	Get(id string) (*Deck, error)
	DrawCards(id string, amount int) ([]card, error)
}
