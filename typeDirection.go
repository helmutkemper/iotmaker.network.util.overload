package iotmakernetworkutiloverload

type Direction int

func (el Direction) String() string {
	return directions[el]
}

var directions = [...]string{
	"in",
	"out",
}
