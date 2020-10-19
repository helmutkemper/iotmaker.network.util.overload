package iotmakernetworkutiloverload

var networks = [...]string{
	"tcp",
	"tcp4",
	"tcp6",
	"unix",
	"unixpacket",
}

func (el TypeNetwork) String() string {
	return networks[el]
}
