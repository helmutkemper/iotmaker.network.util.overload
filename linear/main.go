package main

import (
	"errors"
	"math/rand"
	"net"
	"sync"
	"time"
)

func main() {
	var err error
	//var wg sync.WaitGroup
	//wg.Add(1)
	err = server()
	if err != nil {
		panic(err)
	}
	//wg.Wait()
}

// MinMax (English): Delay between packages
//   Min: Minimal delay
//   Max: Maximal delay
//
// MinMax (Português): Atrasos entre pacotes
//   Min: Atraso mínimo
//   Max: Atraso máximo
type MinMax struct {
	Min time.Duration
	Max time.Duration
}

func (el *MinMax) GenerateTime() (newTime *time.Ticker) {
	seedOfTime := rand.New(rand.NewSource(time.Now().UnixNano()))
	randDuration := time.Duration(seedOfTime.Intn(int(el.Max)-int(el.Min)) + int(el.Min))
	newTime = time.NewTicker(randDuration)

	return
}

const (
	// KDirectionIn (English): active conector to passive receiver
	//
	// KDirectionIn (Português): conector ativo para receptor passivo
	KDirectionIn Direction = iota

	// KDirectionOut (English): passive receiver to active conector
	//
	// KDirectionOut (Português): receptor passivo para conector ativo
	KDirectionOut Direction = iota
)

// Direction (English): Data buffer direction.
//   input: active conector to passive receiver
//   output: passive receiver to active conector
//
// Direction (Português): Direção do buffer de dados
//   input: conector ativo para receptor passivo
//   output: receptor passivo para conector ativo
type Direction int

// directions (English): directions in strings format
//
// directions (Português): direção na forma de string
var directions = [...]string{
	"in",
	"out",
}

// ParserFunc (English): Optional parser function.
//   This function place is between receive and send data package.
//     inData    - Received data from connection
//     inLength  - Received data length
//     direction - Received data direction
//     outData   - Received data after parser
//     outLength - Received data length after parser
//     err       - Error
//
// ParserFunc (Português): Função de parser opcional.
//   Este função fica localizada entre o receptor e o transmissor de pacote de dados.
//     inData    - Dado recebido pela conexão
//     inLength  - Tamanho do dado recebido
//     direction - Direção do dado recebido
//     outData   - Dado recebido depois do parser
//     outLength - Tamanho do dado recebido depois do parser
//     err       - Erro
type ParserFunc func(
	inData []byte,
	inLength int,
	direction Direction,
) (
	outData []byte,
	outLength int,
	err error,
)

// data (English): data package buffer
//
// data (Português): buffer do pacote de dados
type data struct {
	// channel (English): data received
	//
	// channel (Português): dado recebido
	channel chan bool

	// buffer (English): data buffer
	//
	// buffer (Português): buffer do dado
	buffer [][]byte

	m sync.Mutex
}

func (el *data) init() {
	el.buffer = make([][]byte, 0)
	el.channel = make(chan bool, 1)
}

const (
	// KTypeNetworkTcp (English): TCP protocol for IP4 and IP6 network
	//
	// KTypeNetworkTcp (Português): Protocolo TPC para redes com IP4 e IP6
	KTypeNetworkTcp TypeNetwork = iota

	// KTypeNetworkTcp4 (English): TCP protocol for IP4 network only
	//
	// KTypeNetworkTcp4 (Português): Protocolo TPC para redes com apenas o IP4
	KTypeNetworkTcp4

	// KTypeNetworkTcp6 (English): TCP protocol for IP6 network only
	//
	// KTypeNetworkTcp6 (Português): Protocolo TPC para redes com apenas o IP6
	KTypeNetworkTcp6

	// KTypeNetworkUnix (English): UNIX protocol
	//
	// KTypeNetworkUnix (Português): Protocolo UNIX
	KTypeNetworkUnix

	// KTypeNetworkUnixPackage (English): UNIX Package protocol
	//
	// KTypeNetworkUnixPackage (Português): Protocolo Package UNIX
	KTypeNetworkUnixPackage
)

// TypeNetwork (English): Network types. (tcp, tcp4, tcp6, unix and unix package)
//
// TypeNetwork (Português): Tipos de rede. (tcp, tcp4, tcp6, unix and unix package)
type TypeNetwork int

func (el TypeNetwork) String() string {
	return "tcp"
}

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
	//mutex      sync.Mutex
	ticker *time.Ticker
	delays MinMax
}

func (el *TCPConnection) dataConnection(
	conn *net.TCPConn,
	data *data,
	direction Direction,
) {

	go func() {

		var bufferLength int
		var err error

		err = conn.SetNoDelay(false)
		if err != nil {
			el.error <- err
			return
		}
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

			data.m.Lock()
			if cap(data.buffer) == 0 {
				data.buffer = make([][]byte, 0)
			}
			data.buffer = append(data.buffer, buffer[:bufferLength])

			if len(data.channel) == 0 {
				data.channel <- true
			}
			data.m.Unlock()
		}
	}()
}

func server() (err error) {

	var el = &TCPConnection{}
	el.delays.Min = time.Microsecond
	el.delays.Max = time.Microsecond

	var network TypeNetwork
	el.inAddress, err = net.ResolveTCPAddr("tcp", "127.0.0.1:27016")
	if err != nil {
		return
	}

	el.outAddress, err = net.ResolveTCPAddr("tcp", "127.0.0.1:27017")
	if err != nil {
		return
	}

	el.network = network

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

	el.inData.init()
	el.outData.init()

	el.ticker = el.delays.GenerateTime()

	el.listener, err = net.ListenTCP(el.network.String(), el.inAddress)

	for {
		el.inConn, err = el.listener.AcceptTCP()
		go handleOut(el)
		go handleIn(el)
	}
}

func handleIn(el *TCPConnection) {
	var err error
	el.dataConnection(el.inConn, &el.inData, KDirectionIn)

	for {
		select {
		case <-el.inData.channel:
			for {
				if len(el.inData.buffer) == 0 {
					break
				}

				el.inData.m.Lock()
				el.outData.m.Lock()
				_, err = el.outConn.Write(el.inData.buffer[0])
				el.inData.m.Unlock()
				el.outData.m.Unlock()
				if err != nil {
					return
				}

				el.inData.m.Lock()
				el.outData.m.Lock()
				el.inData.buffer = el.inData.buffer[1:]
				el.inData.m.Unlock()
				el.outData.m.Unlock()
			}
		}
	}
}
func handleOut(el *TCPConnection) {
	var err error
	el.outConn, err = net.DialTCP(el.network.String(), nil, el.outAddress)
	if err != nil {
		panic(err)
	}

	el.dataConnection(el.outConn, &el.outData, KDirectionOut)

	defer el.ticker.Stop()

	for {
		select {
		case <-el.ticker.C:
			el.ticker.Stop()

			for {
				if len(el.outData.buffer) == 0 {
					el.ticker = el.delays.GenerateTime()
					break
				}

				el.outData.m.Lock()
				el.inData.m.Lock()
				_, err = el.inConn.Write(el.outData.buffer[0])
				el.outData.m.Unlock()
				el.inData.m.Unlock()
				if err != nil {
					return
				}

				el.outData.m.Lock()
				el.inData.m.Lock()
				el.outData.buffer = el.outData.buffer[1:]
				el.outData.m.Unlock()
				el.inData.m.Unlock()
			}
		}
	}
}
