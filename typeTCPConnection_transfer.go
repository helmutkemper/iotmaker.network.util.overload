package iotmakernetworkutiloverload

import (
	"sync"
)

// transfer (English):
//
// transfer (Português):
func (el *TCPConnection) transfer() (err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

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
	}()
	go func() {
		defer wg.Done()

		for {
			//select {
			//case <-el.ticker.C:
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
			//el.ticker = el.delays.GenerateTime()

		}
		//}
	}()

	wg.Wait()
	return
}
