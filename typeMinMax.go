package iotmakernetworkutiloverload

import (
	"math/rand"
	"time"
)

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

func (el *MinMax) GenerateTime() (newTime *time.Ticker) {
	seedOfTime := rand.New(rand.NewSource(time.Now().UnixNano()))
	randDuration := time.Duration(seedOfTime.Intn(int(el.Max)-int(el.Min)) + int(el.Min))
	newTime = time.NewTicker(randDuration)

	return
}
