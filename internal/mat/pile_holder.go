package mat

type PileHolder int

const (
	FirstPile PileHolder = iota
	SecondPile
	ThirdPile
)

func (h PileHolder) nextPile() PileHolder {
	if h == ThirdPile {
		return FirstPile
	}
	return h + 1
}
