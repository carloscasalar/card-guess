package trick_test

import (
	"fmt"
	"testing"

	"github.com/carloscasalar/card-guess/internal/mat"

	"github.com/stretchr/testify/require"

	"github.com/carloscasalar/card-guess/internal/trick"
	"github.com/stretchr/testify/assert"
)

func Test_NewTrick_should_provide_21_cards(t *testing.T) {
	theTrick, err := trick.New(true)

	require.NoError(t, err)
	assert.Equal(t, 21, theTrick.Sample().Size())
}

func Test_NewTrick_should_provide_21_different_random_cards_each_time(t *testing.T) {
	aTrick, _ := trick.New(true)
	otherTrick, _ := trick.New(true)

	assert.NotEqual(t, aTrick.Sample(), otherTrick.Sample())
}

func Test_NewTrick_should_provide_same_exact_cards_when_dont_shuffle_before_draw(t *testing.T) {
	aTrick, _ := trick.New(false)

	trickCardsStr := aTrick.Sample().String()
	assert.Equal(t, " A[♠]   2[♠]   3[♠]   4[♠]   5[♠]   6[♠]   7[♠]   8[♠]   9[♠]  10[♠]   J[♠]   Q[♠]   K[♠]   A[♥]   2[♥]   3[♥]   4[♥]   5[♥]   6[♥]   7[♥]   8[♥]", trickCardsStr)
}

func Test_NewTrick_should_contain_a_mat_with_initial_sample_split_into_three_piles(t *testing.T) {
	aTrick, _ := trick.New(false)

	assert.Equal(t, " 6[♥]   3[♥]   K[♠]  10[♠]   7[♠]   4[♠]   A[♠]", pileString(aTrick, mat.FirstPile))
	assert.Equal(t, " 7[♥]   4[♥]   A[♥]   J[♠]   8[♠]   5[♠]   2[♠]", pileString(aTrick, mat.SecondPile))
	assert.Equal(t, " 8[♥]   5[♥]   2[♥]   Q[♠]   9[♠]   6[♠]   3[♠]", pileString(aTrick, mat.ThirdPile))
}

func pileString(aTrick trick.Trick, holder mat.PileHolder) string {
	return fmt.Sprintf("%v", aTrick.Mat().Pile(holder))
}
