package iotmakernetworkutiloverload

func (el *TCPConnection) startTicker() {
	el.ticker = el.delays.GenerateTime()
}
