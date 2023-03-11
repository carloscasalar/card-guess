package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/carloscasalar/card-guess/pkg/threepilestrick"

	"github.com/manifoldco/promptui"
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
	trick, err := threepilestrick.New(mustShuffle)
	if err != nil {
		return err
	}

	fmt.Println(cardListString(trick.Sample()))
	fmt.Println("Pick a card from above and hold it in your mind.")
	fmt.Println("Now I'll split the cards into three piles, watch your card.")
	fmt.Println("Then press enter to continue")
	_, _ = fmt.Scanln()

	pileHolder, err := askForThePileWhereTheCardIs(piles(trick))
	if err != nil {
		return err
	}

	fmt.Printf("I've put the pile you choose, the %v, between the other two and split it again into three piles:\n", pileHolder)

	trick, err = trick.MyCardIsInPile(pileHolder)
	if err != nil {
		return err
	}
	pileHolder, err = askForThePileWhereTheCardIs(piles(trick))
	if err != nil {
		return err
	}

	fmt.Printf("For the last time I've put the pile you choose, the %v, between the other two and split it again into three piles:\n", pileHolder)
	trick, err = trick.MyCardIsInPile(pileHolder)
	if err != nil {
		return err
	}
	pileHolder, err = askForThePileWhereTheCardIs(piles(trick))
	if err != nil {
		return err
	}
	trick, err = trick.MyCardIsInPile(pileHolder)
	if err != nil {
		return err
	}

	fmt.Print("Ok, your card is..")
	simulateSuspense()
	guessedCard, err := trick.GuessMyCard()
	if err != nil {
		return err
	}
	fmt.Printf("... %v !\n", *guessedCard)

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

func askForThePileWhereTheCardIs(piles []pileInMat) (threepilestrick.PileHolder, error) {
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

func piles(trick threepilestrick.Trick) []pileInMat {
	return []pileInMat{
		{threepilestrick.FirstPile, trick.FirstPile()},
		{threepilestrick.SecondPile, trick.SecondPile()},
		{threepilestrick.ThirdPile, trick.ThirdPile()},
	}
}

type pileInMat struct {
	Holder threepilestrick.PileHolder
	Pile   threepilestrick.Pile
}

func cardListString(pile threepilestrick.Pile) string {
	cardStrings := make([]string, 0)
	for _, card := range pile {
		cardStrings = append(cardStrings, string(card))
	}
	return strings.Join(cardStrings, " ")
}
