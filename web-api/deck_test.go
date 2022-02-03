package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	assert.Equal(t, 52, deck.RemainingCards())
	assert.Equal(t, "SPADES", deck.Cards[0].suit)
	assert.Equal(t, "ACE", deck.Cards[0].value)
	assert.Equal(t, "HEARTS", deck.Cards[len(deck.Cards)-1].suit)
	assert.Equal(t, "KING", deck.Cards[len(deck.Cards)-1].value)
	assert.False(t, deck.IsShuffled)
}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	deck.Shuffle()

	assert.Equal(t, 52, deck.RemainingCards())
	assert.True(t, deck.IsShuffled)
}

func TestDraw(t *testing.T) {
	deck := newDeck()
	drawnCards, err := deck.Draw(1)

	assert.Equal(t, 51, deck.RemainingCards())
	assert.Nil(t, err)
	assert.Equal(t, "SPADES", drawnCards[0].suit)
	assert.Equal(t, "ACE", drawnCards[0].value)
}

func TestDrawErrorWhenDrawingTooManyCards(t *testing.T) {
	deck := newDeck()
	drawnCards, err := deck.Draw(100)

	assert.Equal(t, 52, deck.RemainingCards())
	assert.NotNil(t, err)
	assert.IsType(t, NotEnoughCardsError(""), err)
	assert.Nil(t, drawnCards)
}
