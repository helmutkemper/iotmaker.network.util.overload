package iotmakernetworkutiloverload

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
	parser     []ParserFunc
	inData     data
	outData    data
	mutex      sync.Mutex
	ticker     *time.Ticker
	delays     MinMax
}
