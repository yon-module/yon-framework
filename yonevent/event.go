package yonevent

import "sync"

var (
	events = make(map[string][]func())
	mu     sync.Mutex
)

func On(event string, fn func()) {
	mu.Lock()
	defer mu.Unlock()
	events[event] = append(events[event], fn)
}

func Emit(event string) {
	mu.Lock()
	defer mu.Unlock()
	if fns, ok := events[event]; ok {
		for _, fn := range fns {
			go fn() // jalan paralel
		}
	}
}
