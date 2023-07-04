package rxgotest

import (
	"log"

	"github.com/reactivex/rxgo/v2"
)

func just() { 
	obs := rxgo.Just(1, 2, 3)()
	for item := range obs.Observe() {
		log.Println(item)
		if item.E != nil {
			panic(item.E)
		}
	}
}