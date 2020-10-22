package iotmakernetworkutiloverload

const (
	// (English): active conector to passive receiver
	//
	// (Português): conector ativo para receptor passivo
	KDirectionIn Direction = iota

	// (English): passive receiver to active conector
	//
	// (Português): receptor passivo para conector ativo
	KDirectionOut Direction = iota
)
