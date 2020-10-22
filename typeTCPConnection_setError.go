package iotmakernetworkutiloverload

// (English):
//
// (Português):
func (el *TCPConnection) setError(err error) {
	el.error <- err
}
