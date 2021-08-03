package iotmakernetworkutiloverload

import (
	"errors"
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
	parser     []ParserFunc
	inData     data
	outData    data
	mutex      sync.Mutex
	ticker     *time.Ticker
	delays     MinMax
}

func (el *TCPConnection) accept() (err error) {
	el.inConn, err = el.listener.AcceptTCP()
	return
}
func (el *TCPConnection) dataConnection(
	conn *net.TCPConn,
	data *data,
	direction Direction,
) {

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
func (el *TCPConnection) dial() (err error) {
	el.outConn, err = net.DialTCP(el.network.String(), nil, el.outAddress)
	return
}
func (el *TCPConnection) inDataConnection() {
	el.dataConnection(el.inConn, &el.inData, KDirectionIn)
}
func (el *TCPConnection) init() {
	el.inData.init()
	el.outData.init()
}
func (el *TCPConnection) listenConn() (err error) {
	el.listener, err = net.ListenTCP(el.network.String(), el.inAddress)
	return
}
func (el *TCPConnection) outDataConnection() {
	el.dataConnection(el.outConn, &el.outData, KDirectionOut)
}
func (el *TCPConnection) ParserAppendTo(fn ParserFunc) {
	if el.parser == nil {
		el.parser = make([]ParserFunc, 0)
	}

	el.parser = append(el.parser, fn)
}
func (el *TCPConnection) ParserReset() {
	el.parser = nil
}
func (el *TCPConnection) SetAddress(network TypeNetwork, inAddress, outAddress string) (err error) {
	el.inAddress, err = net.ResolveTCPAddr(network.String(), inAddress)
	if err != nil {
		return
	}

	el.outAddress, err = net.ResolveTCPAddr(network.String(), outAddress)
	if err != nil {
		return
	}

	el.network = network
	return
}
func (el *TCPConnection) SetDelay(min, max time.Duration) {
	el.delays.Min = min
	el.delays.Max = max
}
func (el *TCPConnection) setError(err error) {
	el.error <- err
}
func (el *TCPConnection) startTicker() {
	el.ticker = el.delays.GenerateTime()
}
func (el *TCPConnection) transfer() (err error) {
	defer el.ticker.Stop()

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
			el.ticker.Stop()

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
