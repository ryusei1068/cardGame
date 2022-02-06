package main

import (
	"fmt"
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

func getCard(card Card) string {
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

func main() {
	fmt.Println("Playing cards")
	deck := generateDeck()
	fmt.Println(deck)
}
