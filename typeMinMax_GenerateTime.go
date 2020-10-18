package iotmaker_network_util_overload

import (
	"math/rand"
	"time"
)

func (el *MinMax) GenerateTime() (newTime *time.Ticker) {
	seedOfTime := rand.New(rand.NewSource(time.Now().UnixNano()))
	randDuration := time.Duration(seedOfTime.Intn(int(el.Max)-int(el.Min)) + int(el.Min))
	newTime = time.NewTicker(randDuration)

	return
}
