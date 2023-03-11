package threepilestrick

import "github.com/carloscasalar/card-guess/internal/mat"

type PileHolder mat.PileHolder

const (
	FirstPile  = PileHolder(mat.FirstPile)
	SecondPile = PileHolder(mat.SecondPile)
	ThirdPile  = PileHolder(mat.ThirdPile)
)

func (h PileHolder) String() string {
	descriptions := map[PileHolder]string{
		FirstPile:  "First pile",
		SecondPile: "Second pile",
		ThirdPile:  "Third pile",
	}
	return descriptions[h]
}
