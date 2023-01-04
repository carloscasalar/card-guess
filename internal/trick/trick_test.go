package trick_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/carloscasalar/card-guess/internal/trick"
	"github.com/stretchr/testify/assert"
)

func Test_NewTrick_should_provide_21_cards(t *testing.T) {
	theTrick, err := trick.New()

	require.NoError(t, err)
	assert.Len(t, theTrick.Cards(), 21)
}
