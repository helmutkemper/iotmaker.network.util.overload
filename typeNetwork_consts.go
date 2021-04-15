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
