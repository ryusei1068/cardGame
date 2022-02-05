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

func main() {
	fmt.Println("Plaing cards")
	card := new_card("2", "♡", 2)
	fmt.Println(card)

}
