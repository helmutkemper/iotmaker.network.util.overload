package iotmakernetworkutiloverload

// (English):
//
// (Português):
func (el *TCPConnection) accept() (err error) {
	el.inConn, err = el.listener.AcceptTCP()
	return
}
