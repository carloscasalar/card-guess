package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	lib "github.com/carloscasalar/card-guess/pkg/threepilestrick"

	"github.com/manifoldco/promptui"

	"github.com/carloscasalar/card-guess/internal/mat"

	"github.com/carloscasalar/card-guess/internal/deck"
)

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
	trick, err := lib.NewDeprecatedTrick(mustShuffle)
	if err != nil {
		return err
	}

	fmt.Println(trick.Sample().String())
	fmt.Println("Pick a card from above and hold it in your mind.")
	fmt.Println("Now I'll split the cards into three piles, watch your card.")
	fmt.Println("Then press enter to continue")
	_, _ = fmt.Scanln()

	pileHolder, err := askForThePileWhereTheCardIs(piles(trick.Mat()))
	if err != nil {
		return err
	}

	fmt.Printf("I've put the pile you choosed, the %v, between the other two and splitted again into three piles:\n", pileHolder)
	// FIXME: this is a smell telling Mat should not be exposed in the trick engine
	sample := trick.Mat().JoinWithPileInTheMiddle(mat.PileHolder(pileHolder))
	theMat, err := splitIntoThreePiles(sample)
	if err != nil {
		return err
	}
	pileHolder, err = askForThePileWhereTheCardIs(piles(theMat))
	if err != nil {
		return err
	}

	fmt.Printf("For the last time I've put the pile you choosed, the %v, between the other two and splitted again into three piles:\n", pileHolder)
	// FIXME: this is a smell telling Mat should not be exposed in the trick engine
	sample = theMat.JoinWithPileInTheMiddle(mat.PileHolder(pileHolder))
	theMat, err = splitIntoThreePiles(sample)
	if err != nil {
		return err
	}
	pileHolder, err = askForThePileWhereTheCardIs(piles(theMat))
	if err != nil {
		return err
	}

	fmt.Print("Ok, your card is..")
	simulateSuspense()
	guessedCard := takeTheFourthCard(theMat, pileHolder)
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

func takeTheFourthCard(theMat lib.Mat, holder lib.PileHolder) lib.DeprecatedCard {
	var pile lib.DeprecatedPile
	switch holder {
	case lib.FirstPile:
		pile = theMat.FirstPile()
	case lib.SecondPile:
		pile = theMat.SecondPile()
	case lib.ThirdPile:
		pile = theMat.ThirdPile()
	}

	var card lib.DeprecatedCard
	const fourth = 4
	for i := 0; i < fourth; i++ {
		card, pile, _ = pile.DrawCard()
	}

	return card
}

func splitIntoThreePiles(sample lib.DeprecatedPile) (lib.Mat, error) {
	theMat := mat.New()
	for {
		var card lib.DeprecatedCard
		var err error
		card, sample, err = sample.DrawCard()
		if err != nil {
			if errors.Is(err, deck.ErrNoMoreCardsInThePile) {
				break
			}
			return nil, err
		}
		theMat = theMat.PlaceIntoNextPile(card)
	}
	return theMat, nil
}

func askForThePileWhereTheCardIs(piles []pileInMat) (lib.PileHolder, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ .DeprecatedPile }}?",
		Active:   "-> {{ .DeprecatedPile | cyan }}",
		Inactive: "   {{ .DeprecatedPile | cyan }}",
		Selected: "{{ .Holder | red | cyan }}, {{ .DeprecatedPile | cyan }}",
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

func piles(aMat lib.Mat) []pileInMat {
	return []pileInMat{
		{lib.FirstPile, aMat.FirstPile()},
		{lib.SecondPile, aMat.SecondPile()},
		{lib.ThirdPile, aMat.ThirdPile()},
	}
}

type pileInMat struct {
	Holder lib.PileHolder
	Pile   lib.DeprecatedPile
}
