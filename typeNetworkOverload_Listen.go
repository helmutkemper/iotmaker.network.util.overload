package iotmaker_network_util_overload

func (el *NetworkOverload) Listen() (err error) {
	el.init()

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
