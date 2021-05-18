package iotmakernetworkutiloverload

import (
	"math/rand"
	"time"
)

// GenerateTime (English): Generate a random time value between max and min
//
// GenerateTime (Português): Gera um valor de tempo aleatório entre máximo e mínimo
func (el *MinMax) GenerateTime() (newTime *time.Ticker) {
	seedOfTime := rand.New(rand.NewSource(time.Now().UnixNano()))
	randDuration := time.Duration(seedOfTime.Intn(int(el.Max)-int(el.Min)) + int(el.Min))
	newTime = time.NewTicker(randDuration)

	return
}
