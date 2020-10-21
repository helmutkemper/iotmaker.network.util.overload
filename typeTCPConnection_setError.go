package iotmakernetworkutiloverload

func (el *TCPConnection) setError(err error) {
	el.error <- err
}
