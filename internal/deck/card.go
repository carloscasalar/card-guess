package deck

import (
	"fmt"
)

type Card interface {
	String() string
}

type cardValue string
type suit string

type card struct {
	value cardValue
	suit  suit
}

func (c card) String() string {
	return fmt.Sprintf("%2s[%s]", c.value, c.suit)
}

func newDeck() Pile {
	values := []cardValue{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []suit{"♠", "♥", "♦", "♣"}

	cards := make([]Card, len(values)*len(suits))
	cardIndex := 0
	for _, suit := range suits {
		for _, value := range values {
			cards[cardIndex] = card{value, suit}
			cardIndex++
		}
	}

	return NewPile(cards...)
}
