package iotmakernetworkutiloverload

import "net"

// (English):
//
// (Português):
func (el *TCPConnection) listenConn() (err error) {
	el.listener, err = net.ListenTCP(el.network.String(), el.inAddress)
	return
}
