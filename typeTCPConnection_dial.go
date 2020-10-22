package iotmakernetworkutiloverload

import "net"

// (English):
//
// (Português):
func (el *TCPConnection) dial() (err error) {
	el.outConn, err = net.DialTCP(el.network.String(), nil, el.outAddress)
	return
}
