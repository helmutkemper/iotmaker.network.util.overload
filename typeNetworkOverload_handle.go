package iotmakernetworkutiloverload

// handle (English): Data connection handle
//
// handle (Português): Gerenciador de dados da conexão
func (el *NetworkOverload) handle() {
	var err error

	err = el.dial()
	if err != nil {
		el.setError(err)
		return
	}

	el.inDataConnection()
	el.outDataConnection()

	err = el.transferInData()
	if err != nil {
		el.setError(err)
		return
	}

	err = el.transferOutData()
	if err != nil {
		el.setError(err)
		return
	}

	return
}
