package deck

import (
	"errors"
	"fmt"
	"strings"
)

type Pile interface {
	DrawCard() (Card, Pile, error)
	AddCard(card Card) Pile
	Cards() []Card
	Size() int
	String() string
}

func NewPile(cards ...Card) Pile {
	if len(cards) == 0 {
		return &emptyPile{}
	}
	firstCard := cards[0]
	otherCards := NewPile(cards[1:]...)
	return &pile{firstCard, otherCards}
}

type pile struct {
	firstCard  Card
	otherCards Pile
}

func (p pile) DrawCard() (Card, Pile, error) {
	newFirstCard, newOtherCards, err := p.otherCards.DrawCard()
	if err != nil {
		if errors.Is(err, ErrNoMoreCardsInThePile) {
			return p.firstCard, emptyPile{}, nil
		}
		return nil, p, err
	}
	drawnCard := p.firstCard
	resultingPile := &pile{newFirstCard, newOtherCards}
	return drawnCard, resultingPile, nil
}

func (p pile) AddCard(card Card) Pile {
	return &pile{
		firstCard:  card,
		otherCards: p,
	}
}

func (p pile) Cards() []Card {
	return append([]Card{p.firstCard}, p.otherCards.Cards()...)
}

func (p pile) Size() int {
	return p.otherCards.Size() + 1
}

func (p pile) String() string {
	var cardStrings = make([]string, p.Size())
	for i, card := range p.Cards() {
		cardStrings[i] = fmt.Sprintf("%v", card)
	}
	return strings.Join(cardStrings, "  ")
}
