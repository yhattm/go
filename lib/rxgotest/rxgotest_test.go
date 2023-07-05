package rxgotest

import (
	"context"
	"testing"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func Test_just(t *testing.T) {
	obs := NewObsFromJust()
	DoOnNext(obs)
	Observe(obs)
}

func Test_Interval(t *testing.T) {
	obs := NewObsFromInterval()
	DoOnNext(obs)
	DoOnCompleted(obs)
	Observe(obs.Take(5))
	time.Sleep(time.Second * 5)
}

func Test_DoOnNext(t *testing.T) {
	obs := NewObsFromInterval()
	ctx, cancel := context.WithCancel(context.Background())
	obs = obs.Take(10, rxgo.WithContext(ctx))
	DoOnNext(obs)
	DoOnCompleted(obs)
	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 5)
}

func Test_FromChannel(t *testing.T) {
	obs := NewObsFromEventSource()
	DoOnCompleted(obs)
	Observe(obs.Take(5))
	time.Sleep(time.Second * 5)
}
