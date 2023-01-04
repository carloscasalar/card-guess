package threepilestrick

type PileHolder int

const (
	FirstPile PileHolder = iota
	SecondPile
	ThirdPile
)

func (h PileHolder) String() string {
	descriptions := map[PileHolder]string{
		FirstPile:  "First pile",
		SecondPile: "Second pile",
		ThirdPile:  "Third pile",
	}
	return descriptions[h]
}
