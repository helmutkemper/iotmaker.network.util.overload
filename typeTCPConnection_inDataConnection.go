package iotmakernetworkutiloverload

func (el *TCPConnection) inDataConnection() {
	el.dataConnection(el.inConn, &el.inData, KDirectionIn)
}
