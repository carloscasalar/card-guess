package deck_test

import (
	"testing"

	"github.com/carloscasalar/card-guess/pkg/threepilestrick"

	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	firstCard   = NewCard("firstCard")
	secondCard  = NewCard("secondCard")
	thirdCard   = NewCard("thirdCard")
	topCard     = NewCard("topCard")
	aMiddleCard = NewCard("aMiddleCard")
	bottomCard  = NewCard("bottomCard")
)

type WhenPileHasNoCardsSuite struct {
	suite.Suite
	pile deck.Pile
}

func Test_when_pile_has_no_cards(t *testing.T) {
	suite.Run(t, new(WhenPileHasNoCardsSuite))
}

func (s *WhenPileHasNoCardsSuite) SetupTest() {
	s.pile = deck.NewPile()
}

func (s *WhenPileHasNoCardsSuite) Test_DrawCard_should_return_error() {
	_, _, err := s.pile.DrawCard()

	assert.ErrorIs(s.T(), err, deck.ErrNoMoreCardsInThePile)
}

func (s *WhenPileHasNoCardsSuite) Test_AddCard_should_add_new_card() {
	pile := s.pile.AddCard(firstCard)

	assert.Equal(s.T(), pile.Cards(), []threepilestrick.Card{firstCard})
}

func (s *WhenPileHasNoCardsSuite) Test_AddCard_twice_should_add_two_cards() {
	pile := s.pile.AddCard(bottomCard)
	pile = pile.AddCard(topCard)

	assert.Equal(s.T(), pile.Cards(), []threepilestrick.Card{topCard, bottomCard})
}

func (s *WhenPileHasNoCardsSuite) Test_Cards_should_return_empty_array() {
	cards := s.pile.Cards()

	assert.Empty(s.T(), cards)
}

func (s *WhenPileHasNoCardsSuite) Test_Size_should_be_zero() {
	assert.Equal(s.T(), 0, s.pile.Size())
}

func (s *WhenPileHasNoCardsSuite) Test_String_representation_should_be_empty() {
	assert.Equal(s.T(), "", s.pile.String())
}

func (s *WhenPileHasNoCardsSuite) Test_StackOnTop_a_pile_should_be_the_stacked_pile_itself() {
	otherPile := deck.NewPile(topCard)

	resultingPile := s.pile.StackOnTopOf(otherPile)

	assert.Equal(s.T(), otherPile, resultingPile)
}

type WhenPileHasOneCardSuite struct {
	suite.Suite
	pile deck.Pile
}

func Test_when_pile_has_one_card(t *testing.T) {
	suite.Run(t, new(WhenPileHasOneCardSuite))
}

func (s *WhenPileHasOneCardSuite) SetupTest() {
	s.pile = deck.NewPile(firstCard)
}

func (s *WhenPileHasOneCardSuite) Test_DrawCard_should_draw_the_card() {
	card, _, err := s.pile.DrawCard()

	require.NoError(s.T(), err)
	assert.Equal(s.T(), card, firstCard)
}

func (s *WhenPileHasOneCardSuite) Test_DrawCard_should_return_error_when_draw_second() {
	_, pile, _ := s.pile.DrawCard()
	_, _, err := pile.DrawCard()

	assert.ErrorIs(s.T(), err, deck.ErrNoMoreCardsInThePile)
}

func (s *WhenPileHasOneCardSuite) Test_Cards_should_list_all_cards_of_the_pile_in_order() {
	cards := s.pile.Cards()

	assert.Equal(s.T(), []threepilestrick.Card{firstCard}, cards)
}

type WhenPileHasTwoCardSuite struct {
	suite.Suite
	pile deck.Pile
}

func Test_when_pile_has_two_card(t *testing.T) {
	suite.Run(t, new(WhenPileHasTwoCardSuite))
}

func (s *WhenPileHasTwoCardSuite) SetupTest() {
	s.pile = deck.NewPile(firstCard, secondCard)
}

func (s *WhenPileHasTwoCardSuite) Test_DrawCard_first_draw_should_draw_first_card() {
	card, _, err := s.pile.DrawCard()

	require.NoError(s.T(), err)
	assert.Equal(s.T(), card, firstCard)
}

func (s *WhenPileHasTwoCardSuite) Test_DrawCard_second_draw_should_draw_second_card() {
	_, pile, _ := s.pile.DrawCard()
	card, _, err := pile.DrawCard()

	require.NoError(s.T(), err)
	assert.Equal(s.T(), card, secondCard)
}

func (s *WhenPileHasTwoCardSuite) Test_DrawCard_third_draw_should_return_error() {
	_, pile, _ := s.pile.DrawCard()
	_, pile, _ = pile.DrawCard()
	_, _, err := pile.DrawCard()

	assert.ErrorIs(s.T(), err, deck.ErrNoMoreCardsInThePile)
}

func (s *WhenPileHasTwoCardSuite) Test_Cards_should_list_all_cards_of_the_pile_in_order() {
	cards := s.pile.Cards()

	assert.Equal(s.T(), []threepilestrick.Card{firstCard, secondCard}, cards)
}

type WhenPileHasThreeCardSuite struct {
	suite.Suite
	pile deck.Pile
}

func Test_when_pile_has_three_card(t *testing.T) {
	suite.Run(t, new(WhenPileHasThreeCardSuite))
}

func (s *WhenPileHasThreeCardSuite) SetupTest() {
	s.pile = deck.NewPile(firstCard, secondCard, thirdCard)
}

func (s *WhenPileHasThreeCardSuite) Test_AddCard_should_add_the_card_on_top_of_the_pile() {
	pile := s.pile.AddCard(topCard)

	assert.Equal(s.T(), []threepilestrick.Card{topCard, firstCard, secondCard, thirdCard}, pile.Cards())
}

func (s *WhenPileHasThreeCardSuite) Test_StackOnTop_should_put_the_whole_pile_on_top() {
	bottomPile := deck.NewPile(aMiddleCard, bottomCard)
	resultingPile := s.pile.StackOnTopOf(bottomPile)

	expectedPile := deck.NewPile(firstCard, secondCard, thirdCard, aMiddleCard, bottomCard)
	assert.Equal(s.T(), expectedPile.String(), resultingPile.String())
}

func (s *WhenPileHasThreeCardSuite) Test_Cards_should_list_all_cards_of_the_pile_in_order() {
	cards := s.pile.Cards()

	assert.Equal(s.T(), []threepilestrick.Card{firstCard, secondCard, thirdCard}, cards)
}

func (s *WhenPileHasThreeCardSuite) Test_String_should_contain_all_cards_of_the_pile_in_order() {
	assert.Equal(s.T(), "firstCard  secondCard  thirdCard", s.pile.String())
}

func NewCard(name string) threepilestrick.Card {
	return &CardStub{name}
}

type CardStub struct {
	name string
}

func (a CardStub) String() string {
	return a.name
}
