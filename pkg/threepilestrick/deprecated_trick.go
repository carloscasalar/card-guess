package threepilestrick

import (
	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/mat"
	"github.com/carloscasalar/card-guess/internal/trick"
)

type DeprecatedTrick interface {
	trick.Trick
}

func NewDeprecatedTrick(shuffleBeforeInitialDraw bool) (DeprecatedTrick, error) {
	return trick.New(shuffleBeforeInitialDraw)
}

// DeprecatedCard represents a card but is now deprecated
type DeprecatedCard interface {
	deck.Card
}

type Mat interface {
	mat.Mat
}

type DeprecatedPile interface {
	deck.Pile
}
