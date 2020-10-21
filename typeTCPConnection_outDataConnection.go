package iotmakernetworkutiloverload

func (el *TCPConnection) outDataConnection() {
	el.dataConnection(el.outConn, &el.outData, KDirectionOut)
}
