package iotmaker_network_util_overload

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
