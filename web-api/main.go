package main

func main() {
	const port string = ":8000"
	deckRepository := NewInMemoryDeckRepository()
	deckService := NewDeckService(deckRepository)
	deckController := NewDeckController(deckService)
	httpRouter := NewMuxRouter()

	httpRouter.Get("/decks/{id}", deckController.GetDeck)
	httpRouter.Post("/decks", deckController.CreateNewDeck)
	httpRouter.Serve((port))
}
