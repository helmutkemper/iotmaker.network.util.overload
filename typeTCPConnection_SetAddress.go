package iotmakernetworkutiloverload

import "net"

// (English):
//
// (Português):
func (el *TCPConnection) SetAddress(network TypeNetwork, inAddress, outAddress string) (err error) {
	el.inAddress, err = net.ResolveTCPAddr(network.String(), inAddress)
	if err != nil {
		return
	}

	el.outAddress, err = net.ResolveTCPAddr(network.String(), outAddress)
	if err != nil {
		return
	}

	el.network = network
	return
}