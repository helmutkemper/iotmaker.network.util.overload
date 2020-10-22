package iotmakernetworkutiloverload

const (
	// (English): TCP protocol for IP4 and IP6 network
	//
	// (Português): Protocolo TPC para redes com IP4 e IP6
	KTypeNetworkTcp TypeNetwork = iota

	// (English): TCP protocol for IP4 network only
	//
	// (Português): Protocolo TPC para redes com apenas o IP4
	KTypeNetworkTcp4

	// (English): TCP protocol for IP6 network only
	//
	// (Português): Protocolo TPC para redes com apenas o IP6
	KTypeNetworkTcp6

	// (English): UNIX protocol
	//
	// (Português): Protocolo UNIX
	KTypeNetworkUnix

	// (English): UNIX Package protocol
	//
	// (Português): Protocolo Package UNIX
	KTypeNetworkUnixPackage
)
