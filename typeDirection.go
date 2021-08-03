package iotmakernetworkutiloverload

const (
	// KDirectionIn (English): active conector to passive receiver
	//
	// KDirectionIn (Português): conector ativo para receptor passivo
	KDirectionIn Direction = iota

	// KDirectionOut (English): passive receiver to active conector
	//
	// KDirectionOut (Português): receptor passivo para conector ativo
	KDirectionOut Direction = iota
)

// Direction (English): Data buffer direction.
//   input: active conector to passive receiver
//   output: passive receiver to active conector
//
// Direction (Português): Direção do buffer de dados
//   input: conector ativo para receptor passivo
//   output: receptor passivo para conector ativo
type Direction int

// directions (English): directions in strings format
//
// directions (Português): direção na forma de string
var directions = [...]string{
	"in",
	"out",
}
