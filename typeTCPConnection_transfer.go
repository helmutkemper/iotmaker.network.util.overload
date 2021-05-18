package iotmakernetworkutiloverload

// transfer (English):
//
// transfer (Português):
func (el *TCPConnection) transferInData() (err error) {
	for {
		select {
		case <-el.inData.channel:
			el.mutex.Lock()
			for {
				if len(el.inData.buffer) == 0 {
					el.mutex.Unlock()
					break
				}

				_, err = el.outConn.Write(el.inData.buffer[0])
				if err != nil {
					return
				}

				el.inData.buffer = el.inData.buffer[1:]
			}
		}
	}
}

func (el *TCPConnection) transferOutData() (err error) {
	for {
		select {
		case <-el.ticker.C:
			el.ticker.Stop()
			el.mutex.Lock()

			for {
				if len(el.outData.buffer) == 0 {
					break
				}

				_, err = el.inConn.Write(el.outData.buffer[0])
				if err != nil {
					return
				}

				el.outData.buffer = el.outData.buffer[1:]
			}

			el.mutex.Unlock()
			el.ticker = el.delays.GenerateTime()
		}
	}
}
