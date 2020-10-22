package iotmakernetworkutiloverload

// (English):
//
// (Português):
func (el *TCPConnection) startTicker() {
	el.ticker = el.delays.GenerateTime()
}
