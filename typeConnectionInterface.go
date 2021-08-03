package iotmakernetworkutiloverload

import "time"

// (English):
//
// (Português):
type ProtocolInterface interface {
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
	ParserReset()
	ParserAppendTo(fn ParserFunc)
}
