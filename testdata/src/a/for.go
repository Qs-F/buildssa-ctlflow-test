package main

import "sync"

func forbreak(b bool) {
	var mu sync.Mutex

	for {
		mu.Lock()
		if b {
			break
		}
		mu.Unlock()
	}
}
