package main

import (
	"log"
	"net/http"
)

//Hardcode in memory store for now
type InMemoryPlayerStore struct{}

//Hardcode in memory store for now
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// Record a win.
func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
