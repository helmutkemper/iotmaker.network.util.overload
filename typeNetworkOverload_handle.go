package iotmakernetworkutiloverload

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
