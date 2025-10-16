package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	votes = make(map[string]int)
	for i := 0; i < 3; i++ {
		go VoteGen()
	}
	time.Sleep(2 * time.Second)
	votesMu.Lock()
	fmt.Println("Результаты:", votes)
	votesMu.Unlock()
}

var (
	votes   map[string]int
	votesMu sync.Mutex
)

func VoteGen() {
	cands := []string{"A", "B", "C"}
	for {
		cand := cands[rand.Intn(len(cands))]
		votesMu.Lock()
		votes[cand]++
		votesMu.Unlock()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
}
