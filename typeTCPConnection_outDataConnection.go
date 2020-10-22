package iotmakernetworkutiloverload

// (English):
//
// (Português):
func (el *TCPConnection) outDataConnection() {
	el.dataConnection(el.outConn, &el.outData, KDirectionOut)
}
