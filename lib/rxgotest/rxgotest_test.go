package rxgotest

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func Test_just(t *testing.T) {
	obs := NewObsFromJust()
	OnNext(obs)
	Observe(obs)
}

func Test_Interval(t *testing.T) {
	obs := NewObsFromInterval()
	OnNext(obs)
	DoOnCompleted(obs)
	Observe(obs.Take(5))
	time.Sleep(time.Second * 5)
}

func Test_OnNext(t *testing.T) {
	obs := NewObsFromInterval()
	ctx, cancel := context.WithCancel(context.Background())
	obs.DoOnNext(
		func(i interface{}) {
			log.Println("OnNext", i)
		},
		rxgo.WithContext(ctx),
	)
	OnNext(obs)
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
