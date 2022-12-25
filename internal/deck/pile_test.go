package deck_test

import (
	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrawCard_should_return_error_when_there_is_no_cards_in_the_pile(t *testing.T) {
	pile := deck.NewPile()

	_, err := pile.DrawCard()

	assert.ErrorIs(t, err, deck.NoMoreCardsInThePile)
}
