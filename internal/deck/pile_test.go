package deck_test

import (
	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDrawCard_should_return_error_when_there_is_no_cards_in_the_pile(t *testing.T) {
	pile := deck.NewPile()

	_, err := pile.DrawCard()

	assert.ErrorIs(t, err, deck.NoMoreCardsInThePile)
}

func TestDrawCard_should_draw_the_card_when_pile_has_only_one_card(t *testing.T) {
	as := NewCard("as")
	pile := deck.NewPile(as)

	card, err := pile.DrawCard()
	require.NoError(t, err)
	assert.Equal(t, *card, as)
}

func TestDrawCard_should_return_error_when_draw_second_card_when_pile_has_only_one_card(t *testing.T) {
	as := NewCard("as")
	pile := deck.NewPile(as)

	_, _ = pile.DrawCard()
	_, err := pile.DrawCard()

	assert.ErrorIs(t, err, deck.NoMoreCardsInThePile)
}

func NewCard(name string) deck.Card {
	return &CardStub{name}
}

type CardStub struct {
	name string
}

func (a CardStub) String() string {
	return a.name
}
