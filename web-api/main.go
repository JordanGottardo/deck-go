package main

func main() {
	const port string = ":8000"
	deckRepository := NewInMemoryDeckRepository()
	deckService := NewDeckService(deckRepository)
	deckController := NewDeckController(deckService)
	httpRouter := NewMuxRouter()

	httpRouter.Post("/decks", deckController.CreateNewDeck)
	httpRouter.Serve((port))
}
