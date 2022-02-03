package main

type Deck struct {
	Id         string `json:"id"`
	isShuffled bool
	cards      []card
}

type card struct {
	suit  string
	value string
}

var suits []string = []string{"Spades, Diamonds, Clubs, Hearts"}
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
