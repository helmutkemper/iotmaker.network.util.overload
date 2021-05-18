package iotmakernetworkutiloverload

// init (English): initializes data buffer
//
// init (Português): inicializa o buffer de dados
func (el *data) init() {
	el.buffer = make([][]byte, 0)
	el.channel = make(chan bool, 1)
	//el.length = make([]int, 0)
}
