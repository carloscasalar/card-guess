package threepilestrick_test

import (
	"strings"
	"testing"

	"github.com/carloscasalar/card-guess/pkg/threepilestrick"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewTrick_should_provide_21_cards(t *testing.T) {
	theTrick, err := threepilestrick.New(true)

	require.NoError(t, err)
	assert.Len(t, theTrick.Sample(), 21)
}

func Test_NewTrick_should_provide_21_different_random_cards_each_time(t *testing.T) {
	aTrick, _ := threepilestrick.New(true)
	otherTrick, _ := threepilestrick.New(true)

	assert.NotEqual(t, aTrick.Sample(), otherTrick.Sample())
}

func Test_NewTrick_should_provide_same_exact_cards_when_dont_shuffle_before_draw(t *testing.T) {
	aTrick, _ := threepilestrick.New(false)

	trickCardsStr := cardsInPile(aTrick.Sample())
	assert.Equal(t, " A[♠]  2[♠]  3[♠]  4[♠]  5[♠]  6[♠]  7[♠]  8[♠]  9[♠] 10[♠]  J[♠]  Q[♠]  K[♠]  A[♥]  2[♥]  3[♥]  4[♥]  5[♥]  6[♥]  7[♥]  8[♥]", trickCardsStr)
}

func Test_on_a_brand_new_trick(t *testing.T) {
	aTrick, _ := threepilestrick.New(false)

	t.Run("the mat should contain three initial piles of cards", func(t *testing.T) {
		assert.Equal(t, " 6[♥]  3[♥]  K[♠] 10[♠]  7[♠]  4[♠]  A[♠]", cardsInPile(aTrick.FirstPile()))
		assert.Equal(t, " 7[♥]  4[♥]  A[♥]  J[♠]  8[♠]  5[♠]  2[♠]", cardsInPile(aTrick.SecondPile()))
		assert.Equal(t, " 8[♥]  5[♥]  2[♥]  Q[♠]  9[♠]  6[♠]  3[♠]", cardsInPile(aTrick.ThirdPile()))
	})

	t.Run("should be impossible to guess the card", func(t *testing.T) {
		_, err := aTrick.GuessMyCard()

		require.Error(t, err)
		assert.Equal(t, "still cannot guess your card, please tell me in which pile is it", err.Error())
	})

}

func cardsInPile(pile threepilestrick.Pile) string {
	var cardStrings []string
	for _, card := range pile {
		cardStrings = append(cardStrings, string(card))
	}
	return strings.Join(cardStrings, " ")
}