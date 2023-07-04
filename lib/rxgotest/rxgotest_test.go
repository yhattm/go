package rxgotest

import (
	"testing"
	"time"
)

func Test_just(t *testing.T) {
	obs := NewObsFromJust()
	OnNext(obs)
	Observe(obs)
}

func Test_Interval(t *testing.T) {
	obs := NewObsFromInterval()
	OnNext(obs)
	Observe(obs.Take(5))
	time.Sleep(time.Second * 5)
}