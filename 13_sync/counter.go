package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// notice this is a pointer
// if we're calling a func on a struc
// we can decide if it's gonna be a pointer
// or a copy if we dont use *
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{}
}
