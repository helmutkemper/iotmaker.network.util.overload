package iotmakernetworkutiloverload

// Listen (English): Network listen
//
// Listen (Português): Ovinte de rede
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
