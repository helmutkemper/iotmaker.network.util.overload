package iotmakernetworkutiloverload

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

	// length (English): buffer length
	//
	// length (Português): tamanho do buffer
	length []int
}
