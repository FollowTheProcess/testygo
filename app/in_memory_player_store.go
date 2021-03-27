package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	// Don't need to take a lock here as concurrent map reads are ok
	return i.store[name]
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.Mutex{}}
}
