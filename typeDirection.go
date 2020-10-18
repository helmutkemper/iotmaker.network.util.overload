package iotmaker_network_util_overload

type Direction int

func (el Direction) String() string {
	return directions[el]
}

var directions = [...]string{
	"in",
	"out",
}
