package iotmakernetworkutiloverload

import "net"

// dataConnection (English): Parses between input package data, custom function and
// package delivery address
//
// dataConnection (Português): Faz o parser entre o dado do pacote de entrada, a função
// customizada e o endereço de entrega do pacote
func (el *TCPConnection) dataConnection(
	conn *net.TCPConn,
	data *data,
	direction Direction,
) {

	go func() {

		var bufferLength int
		var err error
		var buffer = make([]byte, 2048)

		err = conn.SetKeepAlive(true)
		if err != nil {
			el.error <- err
			return
		}

		if data.buffer == nil {
			data.buffer = make([][]byte, 0)
		}

		if data.length == nil {
			data.length = make([]int, 0)
		}

		for {

			bufferLength, err = conn.Read(buffer)
			if err != nil && err.Error() != "EOF" {
				el.error <- err
				return
			}

			if el.parser != nil {
				var err error
				for _, fn := range el.parser {
					buffer, bufferLength, err = fn(buffer, bufferLength, direction)
					if err != nil {
						el.error <- err
						return
					}
				}
			}

			data.buffer = append(data.buffer, buffer[:bufferLength])
			data.length = append(data.length, bufferLength)

			if len(data.channel) == 0 {
				data.channel <- true
			}

			if err != nil && err.Error() == "EOF" {
				break
			}
		}
	}()
}
