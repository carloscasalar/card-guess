package threepilestrick

import "errors"

var ErrStillCannotGuessYourCard = errors.New("still cannot guess your card, please tell me in which pile is it")

var ErrAskMeToGuessTheCardInstead = errors.New("i know which is your card, ask me to guess it instead")
