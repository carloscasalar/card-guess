package mat

import "github.com/carloscasalar/card-guess/pkg/threepilestrick"

type PileHolder threepilestrick.PileHolder

const (
	FirstPile  = PileHolder(threepilestrick.FirstPile)
	SecondPile = PileHolder(threepilestrick.SecondPile)
	ThirdPile  = PileHolder(threepilestrick.ThirdPile)
)

func (h PileHolder) nextPile() PileHolder {
	if h == ThirdPile {
		return FirstPile
	}
	return h + 1
}
