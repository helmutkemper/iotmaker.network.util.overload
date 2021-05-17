package iotmakernetworkutiloverload

// startTicker (English):
//
// startTicker (Português):
func (el *TCPConnection) startTicker() {
	el.ticker.Reset(el.delays.GenerateTime())
}
