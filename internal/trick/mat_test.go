package trick_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/trick"
	"github.com/stretchr/testify/assert"
)

func TestNewMat_should_crate_a_mat_with_three_empty_piles(t *testing.T) {
	mat := trick.NewMat()

	assert.Len(t, mat.GetPile(trick.FirstPile).Cards(), 0)
	assert.Len(t, mat.GetPile(trick.SecondPile).Cards(), 0)
	assert.Len(t, mat.GetPile(trick.ThirdPile).Cards(), 0)
}

func TestMat_PlaceIntoNextPile_should_put_first_card_in_first_pile(t *testing.T) {
	mat := trick.NewMat()

	mat = mat.PlaceIntoNextPile(card("first card"))

	firstPileCards := mat.GetPile(trick.FirstPile).Cards()
	require.Len(t, firstPileCards, 1)
	assert.Equal(t, card("first card"), firstPileCards[0])
}

func TestMat_PlaceIntoNextPile_should_put_second_card_in_second_pile(t *testing.T) {
	mat := newMatWithCards(1)

	mat = mat.PlaceIntoNextPile(card("second card"))

	secondPileCards := mat.GetPile(trick.SecondPile).Cards()
	require.Len(t, secondPileCards, 1)
	assert.Equal(t, card("second card"), secondPileCards[0])
}

func TestMat_PlaceIntoNextPile_should_put_third_card_in_third_pile(t *testing.T) {
	mat := newMatWithCards(2)

	mat = mat.PlaceIntoNextPile(card("third card"))

	thirdPileCards := mat.GetPile(trick.ThirdPile).Cards()
	require.Len(t, thirdPileCards, 1)
	assert.Equal(t, card("third card"), thirdPileCards[0])
}

func TestMat_PlaceIntoNextPile_after_putting_six_cards(t *testing.T) {
	mat := trick.NewMat()

	mat = mat.PlaceIntoNextPile(card("first card")).
		PlaceIntoNextPile(card("second card")).
		PlaceIntoNextPile(card("third card")).
		PlaceIntoNextPile(card("fourth card")).
		PlaceIntoNextPile(card("fifth card")).
		PlaceIntoNextPile(card("sixth card"))

	topCardCases := map[string]struct {
		holder          trick.PileHolder
		expectedTopCard deck.Card
	}{
		"fourth card should be on top of the first pile": {trick.FirstPile, card("fourth card")},
		"fifth card should be on top of the second pile": {trick.SecondPile, card("fifth card")},
		"sixth card should be on top of the third pile":  {trick.ThirdPile, card("sixth card")},
	}
	for sentence, tc := range topCardCases {
		t.Run(sentence, func(t *testing.T) {
			pileCards := mat.GetPile(tc.holder).Cards()
			require.Len(t, pileCards, 2)
			assert.Equal(t, tc.expectedTopCard, pileCards[0])
		})
	}

	bottomCardCases := map[string]struct {
		holder          trick.PileHolder
		expectedTopCard deck.Card
	}{
		"fourth card should be at the bottom of the first pile": {trick.FirstPile, card("first card")},
		"fifth card should be at the bottom of the second pile": {trick.SecondPile, card("second card")},
		"sixth card should be at the bottom of the third pile":  {trick.ThirdPile, card("third card")},
	}
	for sentence, tc := range bottomCardCases {
		t.Run(sentence, func(t *testing.T) {
			pileCards := mat.GetPile(tc.holder).Cards()
			require.Len(t, pileCards, 2)
			assert.Equal(t, tc.expectedTopCard, pileCards[1])
		})
	}
}

func TestMat_PlaceIntoNextPile_should_not_mute_the_mat(t *testing.T) {
	mat := trick.NewMat()

	mat.PlaceIntoNextPile(card("a card"))

	expectedMat := trick.NewMat()
	assert.Equal(t, expectedMat, mat)
}

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

func TestPiles_should_return_piles_in_mat_with_each_holder(t *testing.T) {
	mat := trick.NewMat()

	piles := mat.Piles()

	expectedPiles := []trick.PileInMat{
		{trick.FirstPile, deck.NewPile()},
		{trick.SecondPile, deck.NewPile()},
		{trick.ThirdPile, deck.NewPile()},
	}
	assert.Equal(t, expectedPiles, piles)
}

func newMatWithCards(numberOfCards int) trick.Mat {
	mat := trick.NewMat()
	for i := 0; i < numberOfCards; i++ {
		mat = mat.PlaceIntoNextPile(card(fmt.Sprintf("card %v", i+1)))
	}
	return mat
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
