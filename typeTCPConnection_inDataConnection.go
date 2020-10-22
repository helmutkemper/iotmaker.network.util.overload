package iotmakernetworkutiloverload

// (English):
//
// (Português):
func (el *TCPConnection) inDataConnection() {
	el.dataConnection(el.inConn, &el.inData, KDirectionIn)
}
