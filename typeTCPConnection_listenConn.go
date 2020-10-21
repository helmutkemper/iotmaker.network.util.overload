package iotmakernetworkutiloverload

import "net"

func (el *TCPConnection) listenConn() (err error) {
	el.listener, err = net.ListenTCP(el.network.String(), el.inAddress)
	return
}
