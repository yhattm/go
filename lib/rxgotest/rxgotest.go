package rxgotest

import (
	"log"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func NewObsFromJust() rxgo.Observable {
	return rxgo.Just(1, 2, 3)()
}

func NewObsFromInterval() rxgo.Observable {
	return rxgo.Interval(rxgo.WithDuration(time.Second))
}

func NewObsFromEventSource() rxgo.Observable {
	ch := make(chan rxgo.Item)
	obs := NewObsFromInterval()
	obs.DoOnNext(
		func(i interface{}) {
			ch <- rxgo.Of(i)
		},
	)
	return rxgo.FromEventSource(ch)
}
func Observe(obs rxgo.Observable) {
	for item := range obs.Observe() {
		log.Println(item)
		if item.E != nil {
			panic(item.E)
		}
	}
}

func DoOnNext(obs rxgo.Observable) {
	obs.DoOnNext(
		func(i interface{}) {
			log.Println("DoOnNext", i)
		})

}

func DoOnCompleted(obs rxgo.Observable) {
	obs.DoOnCompleted(
		func() {
			log.Print("DoOnCompleted")
		},
	)
}
