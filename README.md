# deck-go-web-api

Implementation of APIs to model a deck.

# Building and running the project
You can use one of the following to launch the server:
## Executable (on Windows)
You can download the standalone executable included in this repository inside deck-web-api.zip. The executable _deck-web-api.exe_ can be run to launch the server 
## Building and running (all platforms)
You can build and run the server by opening a shell inside the _web-api_ folder and launching
```
go build .
```
An executable will be created inside the same directory. You can then launch it.
## Simply running (all platforms)
You can run the server by opening a shell inside the _web-api_ folder and launching
```
go run .
```

# 

# Launching automated tests
You can launch the automated tests by opening a shell inside the _web-api_ folder and launching (not all code has been tested via automated tests, hoping it gives you evidence of testing knowledge)
```
go test .
```

# Testing the API
The server listens for requests on localhost:8000. You can import the Postman collection (go-deck-web-api.postman_collection.json) in this repository into your local Postman installation to invoke the API.
## Create deck
You can invoke the create deck API at the following endpoint via POST
```
http://localhost:8000/decks?shuffled=true&cards=AS,1D
```
The API accepts shuffled and cards query parameters. If shuffled is not provided, the deck will not be shuffled. If cards is not provided, the full 52-cards deck will be built.
## Open deck
You can invoke the create deck API at the following endpoint via GET
```
http://localhost:8000/decks/{deck_id}
```
## Draw deck
You can invoke the create deck API at the following endpoint via POST
```
http://localhost:8000/decks/{deck_id}/draw?amount=10
```
If the requested amount of cards to draw is higher than the current amount of cards in the deck, an error is returned and no cards will be drawn.
