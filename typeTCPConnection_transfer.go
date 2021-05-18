package iotmakernetworkutiloverload

// transfer (English):
//
// transfer (Português):
func (el *TCPConnection) transfer() (err error) {
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

		case <-el.ticker.C:
			el.mutex.Lock()
			//el.ticker = nil

			for {
				if len(el.outData.buffer) == 0 {
					el.ticker.Reset(el.delays.GenerateTime())
					el.mutex.Unlock()
					break
				}

				_, err = el.inConn.Write(el.outData.buffer[0])
				if err != nil {
					return
				}

				el.outData.buffer = el.outData.buffer[1:]
			}
		}
	}
}
