package iotmakernetworkutiloverload

type ParserFunc func(inData []byte, inLength int, direction Direction) (outData []byte, outLength int, err error)
