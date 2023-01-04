package threepilestrick

import "github.com/carloscasalar/card-guess/internal/trick"

type Trick interface {
	trick.Trick
}

func New(shuffleBeforeInitialDraw bool) (Trick, error) {
	return trick.New(shuffleBeforeInitialDraw)
}
