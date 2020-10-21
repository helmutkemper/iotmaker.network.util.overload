package iotmakernetworkutiloverload

func (el *TCPConnection) accept() (err error) {
	el.inConn, err = el.listener.AcceptTCP()
	return
}
