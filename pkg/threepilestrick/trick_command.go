package threepilestrick

type Command interface {
	Type() CommandType
}
