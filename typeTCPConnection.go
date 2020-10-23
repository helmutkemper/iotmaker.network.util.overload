package iotmakernetworkutiloverload

import (
	"net"
	"sync"
	"time"
)

// TCPConnection (English): Main TCP connection object.
//   This object receives all the necessary functions for the operation of a TCP
//   connection and is separated from the main code so that it can be easily expanded.
//
// TCPConnection (Português): Objeto de conexão TCP.
//   Este objeto recebe todas as funções necessárias para o funcionamento de uma conexão
//   TCP e fica separado do código principal para que o mesmo possa ser expandido de
//   forma fácil.
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
