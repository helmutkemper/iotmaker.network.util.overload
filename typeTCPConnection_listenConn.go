package iotmakernetworkutiloverload

import "net"

// listenConn (English): announces on the local network address
//
// listenConn (Português): anuncia no endereço da rede local
func (el *TCPConnection) listenConn() (err error) {
	el.listener, err = net.ListenTCP(el.network.String(), el.inAddress)
	return
}
