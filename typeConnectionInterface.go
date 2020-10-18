package iotmaker_network_util_overload

import "time"

type ConnectionInterface interface {
	listenConn() (err error)
	accept() (err error)
	inDataConnection()
	outDataConnection()
	transfer() (err error)
	init()
	dial() (err error)
	setError(err error)
	SetAddress(network TypeNetwork, inAddress, outAddress string) (err error)
	SetDelay(min, max time.Duration)
	startTicker()
	verify() (err error)
}
