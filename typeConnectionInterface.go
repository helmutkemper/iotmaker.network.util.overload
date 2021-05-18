package iotmakernetworkutiloverload

import "time"

// ProtocolInterface (English):
//
// ProtocolInterface (Português):
type ProtocolInterface interface {
	// (English):
	//
	// (Português):
	listenConn() (err error)

	// (English):
	//
	// (Português):
	accept() (err error)

	// (English):
	//
	// (Português):
	inDataConnection()

	// (English):
	//
	// (Português):
	outDataConnection()

	// (English):
	//
	// (Português):
	transferInData() (err error)
	transferOutData() (err error)

	// (English):
	//
	// (Português):
	init()

	// (English):
	//
	// (Português):
	dial() (err error)

	// (English):
	//
	// (Português):
	setError(err error)

	// (English):
	//
	// (Português):
	SetAddress(network TypeNetwork, inAddress, outAddress string) (err error)

	// (English):
	//
	// (Português):
	SetDelay(min, max time.Duration)

	// (English):
	//
	// (Português):
	startTicker()

	// (English):
	//
	// (Português):
	verify() (err error)

	// (English):
	//
	// (Português):
	ParserReset()

	// (English):
	//
	// (Português):
	ParserAppendTo(fn ParserFunc)
}
