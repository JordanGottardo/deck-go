package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	Id         string
	IsShuffled bool
	Cards      []card
}

type card struct {
	suit  string
	value string
}

var suits []string = []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
var values []string = []string{"ACE", "1", "2", "3", "4", "5", "6", "7", "8", "9", "JACK", "QUEEN", "KING"}

func GetCard(valueInitial string, suitInitial string) (card, error) {
	var card card

	for _, suit := range suits {
		fmt.Println(suit, suitInitial)
		if strings.HasPrefix(suit, suitInitial) {
			card.suit = suit
			break
		}
	}

	for _, value := range values {
		if strings.HasPrefix(value, valueInitial) {
			card.value = value
			break
		}
	}

	if card.suit == "" || card.value == "" {
		return card, InvalidCardError("Invalid card")
	}

	return card, nil

}

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
		Cards: cards,
	}
}

func newDeckWithRequestedCards(requestedCards []card) Deck {
	return Deck{
		Cards: requestedCards,
	}
}

func (d *Deck) RemainingCards() int {
	return len(d.Cards)
}

func (d *Deck) Shuffle() {
	fmt.Println("Shuffling deck")
	d.IsShuffled = true
	if len(d.Cards) < 2 {
		return
	}

	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)

	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
}

func (d *Deck) Draw(amount int) ([]card, error) {
	fmt.Println("Drawing cards ", amount)

	if amount > d.RemainingCards() {
		return nil, NotEnoughCardsError("Not enough cards")
	}

	drawnCards := d.Cards[:amount]
	d.Cards = d.Cards[amount:]

	fmt.Println(len(d.Cards))

	return drawnCards, nil
}

func (d *Deck) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         string `json:"deck_id"`
		Shuffled   bool   `json:"shuffled"`
		Remanining int    `json:"remaining"`
	}{
		Id:         d.Id,
		Shuffled:   d.IsShuffled,
		Remanining: d.RemainingCards(),
	})
}
