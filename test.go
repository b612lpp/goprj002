package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	q int
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	dd := data{q: 1}

	fmt.Println("start")
	for i := range 10 {

		go co(i, &dd, &wg, &mu)
		wg.Add(1)

	}
	wg.Wait()
}

func co(i int, d *data, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	d.q = d.q + i
	fmt.Println("выполняяется рутина ", i, "значение сейчас", d.q)
	time.Sleep(time.Second * 2)
	mu.Unlock()

}
