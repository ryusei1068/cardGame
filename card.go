package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	value, suit string
	intValue    int32
}

func new_card(value string, suit string, intValue int32) Card {
	var card Card
	card = Card{value: value, suit: suit, intValue: intValue}
	return card
}

func getCardInfo(card Card) string {
	return card.suit + card.value + string(card.intValue)
}

func generateDeck() []Card {
	var deck []Card
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"♡", "♤", "♢", "♧"}

	for _, suit := range suits {
		for i, value := range values {
			intValue := i + 1
			deck = append(deck, new_card(value, suit, int32(intValue)))
		}
	}
	return deck
}

func printDeck(deck []Card) {
	fmt.Println("Displaying cards...")
	for _, card := range deck {
		fmt.Println(getCardInfo(card))
	}
}

func shuffleDeck(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	deckSize := len(deck)
	for i := 0; i < deckSize; i++ {
		randNum := rand.Intn((deckSize - i) + i)
		deck[i], deck[randNum] = deck[randNum], deck[i]
	}
}

func drawOne(deck *[]Card) Card {
	size := len(*deck)
	card := (*deck)[size-1]
	*deck = (*deck)[:size-1]
	return card
}

func main() {
	fmt.Println("Playing cards")
	deck := generateDeck()
	shuffleDeck(deck)
	fmt.Println(deck)
	fmt.Println("Drawed card : ", drawOne(&deck))
	fmt.Println(deck)
}
