package mat_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/mat"
	"github.com/stretchr/testify/assert"
)

func TestNewMat_should_crate_a_mat_with_three_empty_piles(t *testing.T) {
	theMat := mat.New()

	assert.Len(t, theMat.FirstPile().Cards(), 0)
	assert.Len(t, theMat.SecondPile().Cards(), 0)
	assert.Len(t, theMat.ThirdPile().Cards(), 0)
}

func TestMat_PlaceIntoNextPile_should_put_first_card_in_first_pile(t *testing.T) {
	theMat := mat.New()

	theMat = theMat.PlaceIntoNextPile(card("first card"))

	firstPileCards := theMat.FirstPile().Cards()
	require.Len(t, firstPileCards, 1)
	assert.Equal(t, card("first card"), firstPileCards[0])
}

func TestMat_PlaceIntoNextPile_should_put_second_card_in_second_pile(t *testing.T) {
	theMat := newMatWithCards(1)

	theMat = theMat.PlaceIntoNextPile(card("second card"))

	secondPileCards := theMat.SecondPile().Cards()
	require.Len(t, secondPileCards, 1)
	assert.Equal(t, card("second card"), secondPileCards[0])
}

func TestMat_PlaceIntoNextPile_should_put_third_card_in_third_pile(t *testing.T) {
	theMat := newMatWithCards(2)

	theMat = theMat.PlaceIntoNextPile(card("third card"))

	thirdPileCards := theMat.ThirdPile().Cards()
	require.Len(t, thirdPileCards, 1)
	assert.Equal(t, card("third card"), thirdPileCards[0])
}

func TestMat_PlaceIntoNextPile_after_putting_six_cards(t *testing.T) {
	theMat := mat.New()

	theMat = theMat.PlaceIntoNextPile(card("first card")).
		PlaceIntoNextPile(card("second card")).
		PlaceIntoNextPile(card("third card")).
		PlaceIntoNextPile(card("fourth card")).
		PlaceIntoNextPile(card("fifth card")).
		PlaceIntoNextPile(card("sixth card"))

	topCardCases := map[string]struct {
		pile            deck.Pile
		expectedTopCard deck.Card
	}{
		"fourth card should be on top of the first pile": {theMat.FirstPile(), card("fourth card")},
		"fifth card should be on top of the second pile": {theMat.SecondPile(), card("fifth card")},
		"sixth card should be on top of the third pile":  {theMat.ThirdPile(), card("sixth card")},
	}
	for sentence, tc := range topCardCases {
		t.Run(sentence, func(t *testing.T) {
			pileCards := tc.pile.Cards()
			require.Len(t, pileCards, 2)
			assert.Equal(t, tc.expectedTopCard, pileCards[0])
		})
	}

	bottomCardCases := map[string]struct {
		pile            deck.Pile
		expectedTopCard deck.Card
	}{
		"fourth card should be at the bottom of the first pile": {theMat.FirstPile(), card("first card")},
		"fifth card should be at the bottom of the second pile": {theMat.SecondPile(), card("second card")},
		"sixth card should be at the bottom of the third pile":  {theMat.ThirdPile(), card("third card")},
	}
	for sentence, tc := range bottomCardCases {
		t.Run(sentence, func(t *testing.T) {
			pileCards := tc.pile.Cards()
			require.Len(t, pileCards, 2)
			assert.Equal(t, tc.expectedTopCard, pileCards[1])
		})
	}
}

func TestMat_PlaceIntoNextPile_should_not_mute_the_mat(t *testing.T) {
	theMat := mat.New()

	theMat.PlaceIntoNextPile(card("a card"))

	expectedMat := mat.New()
	assert.Equal(t, expectedMat, theMat)
}

func TestMat_JoinWithPileInTheMiddle(t *testing.T) {
	theMat := mat.New().
		PlaceIntoNextPile(card("First Pile, Bottom")).
		PlaceIntoNextPile(card("Second Pile, Bottom")).
		PlaceIntoNextPile(card("Third Pile, Bottom")).
		PlaceIntoNextPile(card("First Pile, Top")).
		PlaceIntoNextPile(card("Second Pile, Top")).
		PlaceIntoNextPile(card("Third Pile, Top"))

	pile := theMat.JoinWithPileInTheMiddle(mat.FirstPile)

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

func newMatWithCards(numberOfCards int) mat.Mat {
	theMat := mat.New()
	for i := 0; i < numberOfCards; i++ {
		theMat = theMat.PlaceIntoNextPile(card(fmt.Sprintf("card %v", i+1)))
	}
	return theMat
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
