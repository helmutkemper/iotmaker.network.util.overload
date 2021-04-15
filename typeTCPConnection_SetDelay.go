package iotmakernetworkutiloverload

import "time"

// SetDelay (English):
//
// SetDelay (Português):
func (el *TCPConnection) SetDelay(min, max time.Duration) {
	el.delays.Min = min
	el.delays.Max = max
}
