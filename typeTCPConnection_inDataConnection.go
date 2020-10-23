package iotmakernetworkutiloverload

// inDataConnection (English): Parses between input package data, custom function and
// package delivery address
//
// inDataConnection (Português): Faz o parser entre o dado do pacote de entrada, a função
// customizada e o endereço de entrega do pacote
func (el *TCPConnection) inDataConnection() {
	el.dataConnection(el.inConn, &el.inData, KDirectionIn)
}
