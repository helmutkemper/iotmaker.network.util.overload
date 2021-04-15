package iotmakernetworkutiloverload

// startTicker (English):
//
// startTicker (Português):
func (el *TCPConnection) startTicker() {
	el.ticker = el.delays.GenerateTime()
}
