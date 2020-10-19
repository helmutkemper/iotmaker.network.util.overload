package iotmakernetworkutiloverload

func (el *NetworkOverload) Listen() (err error) {
	err = el.verify()
	if err != nil {
		return
	}

	el.init()

	err = el.listenConn()
	if err != nil {
		return
	}

	el.startTicker()

	for {
		err = el.accept()
		if err != nil {
			return
		}

		go el.handle()
	}
}
