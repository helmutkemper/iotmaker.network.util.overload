package iotmakernetworkutiloverload

// init (English): initializes the object to can be used
//
// init (Português): inicializa o objeto para ser usado
func (el *TCPConnection) init() {
	el.inData.init()
	el.outData.init()
}
