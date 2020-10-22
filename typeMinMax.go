package iotmakernetworkutiloverload

import "time"

// MinMax (English): Delay between packages
//   Min: Minimal delay
//   Max: Maximal delay
//
// MinMax (Português): Atrasos entre pacotes
//   Min: Atraso mínimo
//   Max: Atraso máximo
type MinMax struct {
	Min time.Duration
	Max time.Duration
}
