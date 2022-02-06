package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	value, suit string
	intValue    int
}

func new_card(value string, suit string, intValue int) Card {
	var card Card
	card = Card{value: value, suit: suit, intValue: intValue}
	return card
}

func getCardInfo(card Card) string {
	return card.suit + card.value
}

func generateDeck() []Card {
	var deck []Card
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"♡", "♤", "♢", "♧"}

	for _, suit := range suits {
		for i, value := range values {
			intValue := i + 1
			deck = append(deck, new_card(value, suit, intValue))
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

type Table struct {
	playersCards [][]Card
	gameMode     string
	deck         []Card
}

func startGame(amountOfPlayers int, gameMode string) Table {
	deck := generateDeck()
	shuffleDeck(deck)
	var playersCards [][]Card

	numberDistribution := initialCards(gameMode)

	for player := 1; player <= amountOfPlayers; player++ {
		var hands []Card
		for i := 1; i <= numberDistribution; i++ {
			hands = append(hands, drawOne(&deck))
		}
		playersCards = append(playersCards, hands)
	}

	table := Table{playersCards: playersCards, gameMode: gameMode, deck: deck}
	return table
}

// default is poker
func initialCards(gameMode string) int {
	if gameMode == "21" {
		return 2
	}
	return 5
}

func getTableInfo(table Table) {
	fmt.Println("Amount of players: ", len(table.playersCards), "...Game mode : ", table.gameMode, ". At this table: ")
	for i, player := range table.playersCards {
		fmt.Println(i+1, "player's cards : ")
		for _, card := range player {
			fmt.Println(getCardInfo(card))
		}
	}
}

func score21Individual(cards []Card) int {
	var score int
	for _, card := range cards {
		score += card.intValue
	}
	if score <= 21 {
		return score
	}
	return 0
}

func winnerOf21(table Table) string {
	var points []int
	cache := make(map[int]int)

	for _, cards := range table.playersCards {
		point := score21Individual(cards)
		points = append(points, point)
		_, exist := cache[point]
		if exist {
			cache[point] += 1
		} else {
			cache[point] = 1
		}
	}
	fmt.Println(points)
	winnerIndex := maxInArrayIndex(points)
	if cache[points[winnerIndex]] > 1 {
		return "It is a draw"
	} else if cache[points[winnerIndex]] >= 0 {
		winner := winnerIndex + 1
		result := "player " + string(winner) + " is the winner"
		return result
	}
	return "No winners..."
}

func maxInArrayIndex(arr []int) int {
	maxIndex := 0
	maxValue := arr[0]

	for i, num := range arr {
		if num > maxValue {
			maxValue = num
			maxIndex = i
		}
	}

	return maxIndex
}

func main() {
	fmt.Println("Playing cards")
	table := startGame(4, "21")
	getTableInfo(table)
	fmt.Println(winnerOf21(table))
}
