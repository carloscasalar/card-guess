package trick_test

import (
	"testing"

	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/trick"
	"github.com/stretchr/testify/assert"
)

func TestMat_JoinWithPileInTheMiddle(t *testing.T) {
	mat := trick.NewMat().
		PlaceIntoNextPile(card("First Pile, Bottom")).
		PlaceIntoNextPile(card("Second Pile, Bottom")).
		PlaceIntoNextPile(card("Third Pile, Bottom")).
		PlaceIntoNextPile(card("First Pile, Top")).
		PlaceIntoNextPile(card("Second Pile, Top")).
		PlaceIntoNextPile(card("Third Pile, Top"))

	pile := mat.JoinWithPileInTheMiddle(trick.FirstPile)

	expectedPile := deck.NewPile(
		card("Second Pile, Top"),
		card("Second Pile, Bottom"),
		card("First Pile, Top"),
		card("First Pile, Bottom"),
		card("Third Pile, Top"),
		card("Third Pile, Bottom"),
	)

	assert.Equal(t, expectedPile.String(), pile.String())
}

func card(name string) deck.Card {
	return &CardStub{name}
}

type CardStub struct {
	name string
}

func (a CardStub) String() string {
	return a.name
}
