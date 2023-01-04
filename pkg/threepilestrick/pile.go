package threepilestrick

type Pile interface {
	DrawCard() (Card, Pile, error)
	AddCard(card Card) Pile
	StackOnTopOf(Pile) Pile
	Cards() []Card
	Size() int
	String() string
}
