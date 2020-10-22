package iotmakernetworkutiloverload

import "net"

// (English):
//
// (Português):
func (el *TCPConnection) dataConnection(conn *net.TCPConn, data *data, direction Direction) {
	go func() {

		var bufferLength int
		var err error

		err = conn.SetKeepAlive(true)
		if err != nil {
			el.error <- err
			return
		}

		for {
			var buffer = make([]byte, 2048)
			bufferLength, err = conn.Read(buffer)
			if err != nil && err.Error() != "EOF" {
				el.error <- err
				return
			}

			if err != nil && err.Error() == "EOF" {
				break
			}

			if el.parser != nil {
				for _, fn := range el.parser {
					buffer, bufferLength, err = fn(buffer, bufferLength, direction)
					if err != nil {
						el.error <- err
						return
					}
				}
			}

			if cap(data.buffer) == 0 {
				data.buffer = make([][]byte, 0)
			}
			data.buffer = append(data.buffer, buffer[:bufferLength])

			if cap(data.length) == 0 {
				data.length = make([]int, 0)
			}
			data.length = append(data.length, bufferLength)

			if len(data.channel) == 0 {
				data.channel <- true
			}
		}
	}()
}
