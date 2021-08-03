package iotmakernetworkutiloverload

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

type NetworkOverload struct {

	// ProtocolInterface (English): Protocol interface (interface code contract)
	//
	// ProtocolInterface (Português): Interface do protocolo (interface do protocolo de
	// código)
	ProtocolInterface
}

func (el *NetworkOverload) handle() {
	var err error

	err = el.dial()
	if err != nil {
		el.setError(err)
		return
	}

	el.inDataConnection()
	el.outDataConnection()

	err = el.transfer()
	if err != nil {
		el.setError(err)
		return
	}

	return
}

func (el *NetworkOverload) Listen() (err error) {
	err = el.verify()
	if err != nil {
		return
	}

	el.init()
	el.startTicker()

	err = el.listenConn()
	if err != nil {
		return
	}

	for {
		err = el.accept()
		if err != nil {
			return
		}

		go el.handle()
	}
}
