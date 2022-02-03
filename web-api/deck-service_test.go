package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (mock mockRepository) Save(deck *Deck) (*Deck, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.(*Deck), args.Error(1)
}

func (mock mockRepository) Get(id string) (*Deck, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.(*Deck), args.Error(1)
}

func (mock mockRepository) DrawCards(id string, amount int) ([]card, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.([]card), args.Error(1)
}

func TestGet(t *testing.T) {
	mockRepository := new(mockRepository)
	expectedDeck := ADeck()
	mockRepository.On("Get").Return(expectedDeck, nil)
	service := NewDeckService(mockRepository)

	result, err := service.Get("123")

	mockRepository.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedDeck.Id, result.Id)
	assert.Equal(t, expectedDeck.IsShuffled, result.IsShuffled)
	assert.Equal(t, expectedDeck.RemainingCards(), result.RemainingCards())
	assert.Equal(t, expectedDeck.Cards[0].suit, result.Cards[0].suit)
	assert.Equal(t, expectedDeck.Cards[0].value, result.Cards[0].value)
}

func TestCreateFullDeck(t *testing.T) {
	mockRepository := new(mockRepository)
	createDeckDto := CreateDeckDto{
		Shuffled: false,
	}
	service := NewDeckService(mockRepository)
	expectedDeck := ADeck()
	mockRepository.On("Save").Return(expectedDeck, nil)

	result, err := service.Create(createDeckDto)

	mockRepository.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedDeck.RemainingCards(), result.RemainingCards())
	assert.Equal(t, expectedDeck.IsShuffled, result.IsShuffled)
	assert.Equal(t, expectedDeck.Cards[0].suit, result.Cards[0].suit)
	assert.Equal(t, expectedDeck.Cards[0].value, result.Cards[0].value)
}

func ADeck() *Deck {
	expectedCard := card{
		suit:  "aSuit",
		value: "aValue",
	}
	return &Deck{
		Id:         "123",
		IsShuffled: true,
		Cards:      []card{expectedCard},
	}
}
