package iotmakernetworkutiloverload

// accept (English): accepts the next incoming call and mounts the new connection.
//
// accept (Português): aceita a próxima chamada e monta a nova conexão
func (el *TCPConnection) accept() (err error) {
	el.inConn, err = el.listener.AcceptTCP()
	return
}
