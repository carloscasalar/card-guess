package threepilestrick

type Trick interface {
	Sample() Pile
	Mat() Mat
}
