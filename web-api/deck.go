package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Id         string
	IsShuffled bool
	cards      []card
}

type card struct {
	suit  string
	value string
}

var suits []string = []string{"Spades", "Diamonds", "Clubs", "Hearts"}
var values []string = []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Jack", "Queen", "King"}

func newDeck() Deck {
	var cards []card = make([]card, 0)

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			card := card{
				suit:  suits[i],
				value: values[j],
			}

			cards = append(cards, card)
		}
	}

	return Deck{
		cards: cards,
	}
}

func (d *Deck) remainingCards() int {
	return len(d.cards)
}

func (d *Deck) Shuffle() {
	fmt.Println("Shuffling deck")
	d.IsShuffled = true
	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)

	for i := range d.cards {
		newPosition := r.Intn(len(d.cards) - 1)

		d.cards[i], d.cards[newPosition] = d.cards[newPosition], d.cards[i]
	}
}

func (d *Deck) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         string `json:"id"`
		Shuffled   bool   `json:"shuffled"`
		Remanining int    `json:"remaining"`
	}{
		Id:         d.Id,
		Shuffled:   d.IsShuffled,
		Remanining: d.remainingCards(),
	})
}
