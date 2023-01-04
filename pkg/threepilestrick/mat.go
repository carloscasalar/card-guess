package threepilestrick

type Mat interface {
	PlaceIntoNextPile(card Card) Mat
	JoinWithPileInTheMiddle(holder PileHolder) Pile
	FirstPile() Pile
	SecondPile() Pile
	ThirdPile() Pile
}
