package iotmakernetworkutiloverload

// setError (English):
//
// setError (Português):
func (el *TCPConnection) setError(err error) {
	el.error <- err
}
