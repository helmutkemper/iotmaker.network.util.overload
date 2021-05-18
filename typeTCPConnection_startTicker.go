package iotmakernetworkutiloverload

import (
	"time"
)

// startTicker (English):
//
// startTicker (Português):
func (el *TCPConnection) startTicker() {
	el.ticker = time.NewTicker(el.delays.GenerateTime())
	//el.ticker.Reset(el.delays.GenerateTime())
}
