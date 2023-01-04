package deck_test

import (
	"testing"

	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const StandardDeckSize = 52

type DealerSuite struct {
	suite.Suite
	dealer deck.Dealer
}

func Test_dealer(t *testing.T) {
	suite.Run(t, new(DealerSuite))
}

func (s *DealerSuite) SetupTest() {
	s.dealer = deck.NewDealer()
}

func (s *DealerSuite) Test_deal_should_deal_a_card() {
	card, err := s.dealer.Deal()

	require.NoError(s.T(), err)
	assert.NotNil(s.T(), card)
}

func (s *DealerSuite) Test_deal_should_deal_up_52() {
	_, err := dealCards(s.dealer, StandardDeckSize)

	assert.NoError(s.T(), err)
}

func (s *DealerSuite) Test_all_card_dealt_should_be_different() {
	cardsDealt, _ := dealCards(s.dealer, StandardDeckSize)

	uniqueCardsDealt := toMapOfCards(cardsDealt)

	assert.Len(s.T(), uniqueCardsDealt, StandardDeckSize)
}

func (s *DealerSuite) Test_trying_to_deal_more_than_52_cards_should_return_error() {
	_, err := dealCards(s.dealer, 53)

	require.Error(s.T(), err)
}

func (s *DealerSuite) Test_shuffled_cards_should_contain_different_in_different_order() {
	shuffleDealer := deck.NewDealer()
	shuffleDealer.ShuffleCards()
	shuffledCards, _ := dealCards(shuffleDealer, StandardDeckSize)
	nonShuffledCards, _ := dealCards(s.dealer, StandardDeckSize)

	assert.Equal(s.T(), len(nonShuffledCards), len(shuffledCards))
	assert.NotEqual(s.T(), nonShuffledCards, shuffledCards)
	assert.Equal(s.T(), toMapOfCards(nonShuffledCards), toMapOfCards(shuffledCards))
}

func toMapOfCards(cardsDealt []deck.Card) map[string]bool {
	uniqueCardsDealt := map[string]bool{}
	for _, card := range cardsDealt {
		cardStr := card.String()
		uniqueCardsDealt[cardStr] = true
	}
	return uniqueCardsDealt
}

func dealCards(dealer deck.Dealer, numberOfCards int) ([]deck.Card, error) {
	cardsDealt := make([]deck.Card, numberOfCards)
	for i := 0; i < numberOfCards; i++ {
		card, err := dealer.Deal()
		if err != nil {
			return nil, err
		}
		cardsDealt[i] = card
	}
	return cardsDealt, nil
}
