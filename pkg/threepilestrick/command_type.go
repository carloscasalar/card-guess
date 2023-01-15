package threepilestrick

type CommandType int

const (
	ChoosePileWhereYourCardIs CommandType = iota
	GuessMyCard
)
