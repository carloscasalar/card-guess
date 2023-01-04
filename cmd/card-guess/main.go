package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"

	"github.com/carloscasalar/card-guess/internal/mat"

	"github.com/carloscasalar/card-guess/internal/deck"
)

const TrickSampleSize = 21

func main() {
	mustShuffle := flag.Bool("shuffle-before-initial-sample", true, "Tells if you want to shuffle before drawing the initial set of cards")
	flag.Parse()

	if mustShuffle == nil {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := run(*mustShuffle); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run(mustShuffle bool) error {
	dealer := deck.NewDealer()
	if mustShuffle {
		dealer.ShuffleCards()
	}

	sample := deck.NewPile()
	for i := 0; i < TrickSampleSize; i++ {
		card, err := dealer.Deal()
		if err != nil {
			return fmt.Errorf("unexpected error while dealing the card %vth: %w", i+1, err)
		}
		sample = sample.AddCard(card)
	}

	fmt.Println(sample.String())
	fmt.Println("Pick a card from above and hold it in your mind.")
	fmt.Println("Now I'll split the cards into three piles, watch your card.")
	fmt.Println("Then press enter to continue")
	_, _ = fmt.Scanln()

	mat, err := splitIntoThreePiles(sample)
	if err != nil {
		return err
	}

	pileHolder, err := askForThePileWhereTheCardIs(piles(mat))
	if err != nil {
		return err
	}

	fmt.Printf("I've put the pile you choosed, the %v, between the other two and splitted again into three piles:\n", pileHolder)
	sample = mat.JoinWithPileInTheMiddle(pileHolder)
	mat, err = splitIntoThreePiles(sample)
	if err != nil {
		return err
	}
	pileHolder, err = askForThePileWhereTheCardIs(piles(mat))
	if err != nil {
		return err
	}

	fmt.Printf("For the last time I've put the pile you choosed, the %v, between the other two and splitted again into three piles:\n", pileHolder)
	sample = mat.JoinWithPileInTheMiddle(pileHolder)
	mat, err = splitIntoThreePiles(sample)
	if err != nil {
		return err
	}
	pileHolder, err = askForThePileWhereTheCardIs(piles(mat))
	if err != nil {
		return err
	}

	fmt.Print("Ok, your card is..")
	simulateSuspense()
	guessedCard := takeTheFourthCard(mat, pileHolder)
	fmt.Printf("... %v !\n", guessedCard)

	return nil
}

func simulateSuspense() {
	const suspenseTime = 1 * time.Second
	time.Sleep(suspenseTime)
	fmt.Print(".")
	time.Sleep(suspenseTime)
	fmt.Print(".")
	time.Sleep(suspenseTime)
}

func takeTheFourthCard(theMat mat.Mat, holder mat.PileHolder) deck.Card {
	var pile deck.Pile
	switch holder {
	case mat.FirstPile:
		pile = theMat.FirstPile()
	case mat.SecondPile:
		pile = theMat.SecondPile()
	case mat.ThirdPile:
		pile = theMat.ThirdPile()
	}

	var card deck.Card
	const fourth = 4
	for i := 0; i < fourth; i++ {
		card, pile, _ = pile.DrawCard()
	}

	return card
}

func splitIntoThreePiles(sample deck.Pile) (mat.Mat, error) {
	mat := mat.New()
	for {
		var card deck.Card
		var err error
		card, sample, err = sample.DrawCard()
		if err != nil {
			if errors.Is(err, deck.ErrNoMoreCardsInThePile) {
				break
			}
			return nil, err
		}
		mat = mat.PlaceIntoNextPile(card)
	}
	return mat, nil
}

func askForThePileWhereTheCardIs(piles []pileInMat) (mat.PileHolder, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Pile }}?",
		Active:   "-> {{ .Pile | cyan }}",
		Inactive: "   {{ .Pile | cyan }}",
		Selected: "{{ .Holder | red | cyan }}, {{ .Pile | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select the pile where your card is",
		Items:     piles,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return -1, fmt.Errorf("failed to retrieve your chosen pile %w", err)
	}

	return piles[i].Holder, nil
}

func piles(aMat mat.Mat) []pileInMat {
	return []pileInMat{
		{mat.FirstPile, aMat.FirstPile()},
		{mat.SecondPile, aMat.SecondPile()},
		{mat.ThirdPile, aMat.ThirdPile()},
	}
}

type pileInMat struct {
	Holder mat.PileHolder
	Pile   deck.Pile
}
