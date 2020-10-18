package iotmaker_network_util_overload

import (
	"net"
	"sync"
	"time"
)

type TCPConnection struct {
	network    TypeNetwork
	inAddress  *net.TCPAddr
	outAddress *net.TCPAddr
	listener   *net.TCPListener
	inConn     *net.TCPConn
	outConn    *net.TCPConn
	error      chan error
	parser     []func(inData []byte, inLength int, direction direction) (outData []byte, outLength int, err error)
	inData     data
	outData    data
	mutex      sync.Mutex
	ticker     *time.Ticker
	delays     MinMax
}

func (el *TCPConnection) init() {
	el.inData.init()
	el.outData.init()
}

func (el *TCPConnection) setError(err error) {
	el.error <- err
}

func (el *TCPConnection) listenConn() (err error) {
	el.listener, err = net.ListenTCP(el.network.String(), el.inAddress)
	return
}

func (el *TCPConnection) dial() (err error) {
	el.outConn, err = net.DialTCP(el.network.String(), nil, el.outAddress)
	return
}

func (el *TCPConnection) accept() (err error) {
	el.inConn, err = el.listener.AcceptTCP()
	return
}

func (el *TCPConnection) inDataConnection() {
	el.dataConnection(el.inConn, &el.inData, KDirectionIn)
}

func (el *TCPConnection) outDataConnection() {
	el.dataConnection(el.outConn, &el.outData, KDirectionOut)
}

func (el *TCPConnection) dataConnection(conn *net.TCPConn, data *data, direction direction) {
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
			el.ticker = nil

			for {
				if len(el.outData.buffer) == 0 {
					el.ticker = el.delays.GenerateTime()
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
