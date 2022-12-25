package deck_test

import (
	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	firstCard  = NewCard("firstCard")
	secondCard = NewCard("secondCard")
	thirdCard  = NewCard("thirdCard")
)

type WhenPileHasNoCardsSuite struct {
	suite.Suite
	pile *deck.Pile
}

func Test_when_pile_has_no_cards(t *testing.T) {
	suite.Run(t, new(WhenPileHasNoCardsSuite))
}

func (s *WhenPileHasNoCardsSuite) SetupTest() {
	s.pile = deck.NewPile()
}

func (s *WhenPileHasNoCardsSuite) Test_DrawCard_should_return_error() {
	_, err := s.pile.DrawCard()

	assert.ErrorIs(s.T(), err, deck.NoMoreCardsInThePile)
}

func (s *WhenPileHasNoCardsSuite) Test_Cards_should_return_empty_array() {
	cards := s.pile.Cards()

	assert.Empty(s.T(), cards)
}

type WhenPileHasOneCardSuite struct {
	suite.Suite
	pile *deck.Pile
}

func Test_when_pile_has_one_card(t *testing.T) {
	suite.Run(t, new(WhenPileHasOneCardSuite))
}

func (s *WhenPileHasOneCardSuite) SetupTest() {
	s.pile = deck.NewPile(firstCard)
}

func (s *WhenPileHasOneCardSuite) Test_DrawCard_should_draw_the_card() {
	card, err := s.pile.DrawCard()

	require.NoError(s.T(), err)
	assert.Equal(s.T(), *card, firstCard)
}

func (s *WhenPileHasOneCardSuite) Test_DrawCard_should_return_error_when_draw_second() {
	_, _ = s.pile.DrawCard()
	_, err := s.pile.DrawCard()

	assert.ErrorIs(s.T(), err, deck.NoMoreCardsInThePile)
}

func (s *WhenPileHasOneCardSuite) Test_Cards_should_list_all_cards_of_the_pile_in_order() {
	cards := s.pile.Cards()

	assert.Equal(s.T(), []deck.Card{firstCard}, cards)
}

type WhenPileHasTwoCardSuite struct {
	suite.Suite
	pile *deck.Pile
}

func Test_when_pile_has_two_card(t *testing.T) {
	suite.Run(t, new(WhenPileHasTwoCardSuite))
}

func (s *WhenPileHasTwoCardSuite) SetupTest() {
	s.pile = deck.NewPile(firstCard, secondCard)
}

func (s *WhenPileHasTwoCardSuite) Test_DrawCard_first_draw_should_draw_first_card() {
	card, err := s.pile.DrawCard()

	require.NoError(s.T(), err)
	assert.Equal(s.T(), *card, firstCard)
}

func (s *WhenPileHasTwoCardSuite) Test_DrawCard_second_draw_should_draw_second_card() {
	_, _ = s.pile.DrawCard()
	card, err := s.pile.DrawCard()

	require.NoError(s.T(), err)
	assert.Equal(s.T(), *card, secondCard)
}

func (s *WhenPileHasTwoCardSuite) Test_DrawCard_third_draw_should_return_error() {
	_, _ = s.pile.DrawCard()
	_, _ = s.pile.DrawCard()
	_, err := s.pile.DrawCard()

	assert.ErrorIs(s.T(), err, deck.NoMoreCardsInThePile)
}

func (s *WhenPileHasTwoCardSuite) Test_Cards_should_list_all_cards_of_the_pile_in_order() {
	cards := s.pile.Cards()

	assert.Equal(s.T(), []deck.Card{firstCard, secondCard}, cards)
}

type WhenPileHasThreeCardSuite struct {
	suite.Suite
	pile *deck.Pile
}

func Test_when_pile_has_three_card(t *testing.T) {
	suite.Run(t, new(WhenPileHasThreeCardSuite))
}

func (s *WhenPileHasThreeCardSuite) SetupTest() {
	s.pile = deck.NewPile(firstCard, secondCard, thirdCard)
}

func (s *WhenPileHasThreeCardSuite) Test_Cards_should_list_all_cards_of_the_pile_in_order() {
	cards := s.pile.Cards()

	assert.Equal(s.T(), []deck.Card{firstCard, secondCard, thirdCard}, cards)
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
