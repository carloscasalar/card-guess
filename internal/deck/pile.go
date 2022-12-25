package deck

import (
	"fmt"
	"strings"
)

type Pile interface {
	DrawCard() (Card, Pile, error)
	AddCard(card Card) Pile
	Cards() []Card
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
		if err == ErrNoMoreCardsInThePile {
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

func (p pile) String() string {
	var cardStrings []string
	for _, card := range p.Cards() {
		cardStrings = append(cardStrings, fmt.Sprintf("%v", card))
	}
	return strings.Join(cardStrings, "  ")
}
