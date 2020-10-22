package iotmakernetworkutiloverload

import "errors"

// (English):
//
// (Português):
func (el *TCPConnection) verify() (err error) {
	if el.delays.Min == 0 {
		err = errors.New("please, set min and max timers")
		return
	}

	if el.delays.Max == 0 {
		err = errors.New("please, set min and max timers")
		return
	}

	if el.delays.Min == el.delays.Max {
		el.delays.Max += 1
	}

	return
}
