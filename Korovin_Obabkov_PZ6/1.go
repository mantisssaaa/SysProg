package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu  sync.Mutex
		wg  sync.WaitGroup
		num int
	)
	wg.Add(3)
	One := func() {
		defer wg.Done()
		mu.Lock()
		for i := 0; i < 10; i++ {
			num++
			fmt.Println(num)
		}
		mu.Unlock()
	}
	Two := func() {
		defer wg.Done()
		mu.Lock()
		for i := 0; i < 5; i++ {
			num += 2
			fmt.Println(num)
		}
		mu.Unlock()
	}
	Three := func() {
		defer wg.Done()
		mu.Lock()
		for i := 0; i < 5; i++ {
			num += 3
			fmt.Println(num)
		}
		mu.Unlock()
	}

	fmt.Println("Счетчик:")
	go One()
	go Two()
	go Three()
	wg.Wait()
}
