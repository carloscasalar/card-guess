package trick

type holder int

const (
	firstPile holder = iota
	secondPile
	thirdPile
)

func (h holder) NextPile() holder {
	if h == thirdPile {
		return firstPile
	}
	return h + 1
}
