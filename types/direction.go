package types

type Direction int

const (
	East Direction = iota
	West
	South
	North
)

var DirectionArr [4]Direction = [4]Direction{East, West, South, North}
