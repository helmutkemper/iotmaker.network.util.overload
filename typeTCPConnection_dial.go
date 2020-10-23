package iotmakernetworkutiloverload

import "net"

// dial (English): connects to the address on the named network.
//
// dial (Português): conecta o endereço da rede designada.
func (el *TCPConnection) dial() (err error) {
	el.outConn, err = net.DialTCP(el.network.String(), nil, el.outAddress)
	return
}
