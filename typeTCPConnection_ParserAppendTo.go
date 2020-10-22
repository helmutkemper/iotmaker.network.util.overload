package iotmakernetworkutiloverload

// (English):
//
// (Português):
func (el *TCPConnection) ParserAppendTo(fn ParserFunc) {
	if el.parser == nil {
		el.parser = make([]ParserFunc, 0)
	}

	el.parser = append(el.parser, fn)
}
