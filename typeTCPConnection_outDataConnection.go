package iotmakernetworkutiloverload

// outDataConnection (English): Parses between input package data, custom function and
// package delivery address
//
// outDataConnection (Português): Faz o parser entre o dado do pacote de entrada, a função
// customizada e o endereço de entrega do pacote
func (el *TCPConnection) outDataConnection() {
	el.dataConnection(el.outConn, &el.outData, KDirectionOut)
}
