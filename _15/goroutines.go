package main

import (
	"fmt"
	"sync"
)

// Goroutinrs;- it helps to run tasks simultanuously
//  basically it does multi threading

// Waitgroups:- waitgroups helps to manage goroutiines to execute
//  completely before exiting a program

func task(id int, w *sync.WaitGroup) {
	defer w.Done() //defer keyword used to execute the line at last
	fmt.Println("executing ", id)
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}
