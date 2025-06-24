package main

import (
	"fmt"
	"sync"
)

// Mutex:- It help to lock the resoure being modified by concuureenly
// to avoid race condition

type myst struct {
	num int
	mt  sync.Mutex
}

func (m *myst) increase(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		m.mt.Unlock()
	}()

	m.mt.Lock()
	m.num += 1

}

func main() {
	var wg sync.WaitGroup
	mystint := myst{
		num: 0,
	}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go mystint.increase(&wg)
	}
	wg.Wait()
	fmt.Println(mystint.num)
}
