package main

type DeckNotFoundError string

func (e DeckNotFoundError) Error() string {
	return string(e)
}

type NotEnoughCardsError string

func (e NotEnoughCardsError) Error() string {
	return string(e)
}

type InvalidCardError string

func (e InvalidCardError) Error() string {
	return string(e)
}

type ServiceError struct {
	Message string `json:"message"`
}
