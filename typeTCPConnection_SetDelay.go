package iotmakernetworkutiloverload

import "time"

func (el *TCPConnection) SetDelay(min, max time.Duration) {
	el.delays.Min = min
	el.delays.Max = max
}
