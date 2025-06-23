package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channels:-
// Channels are used to send data between the go routines

func chanReceiver(nums chan int) {
	for i := range nums {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {

	mychan := make(chan int)

	go chanReceiver(mychan)

	for {
		mychan <- rand.Intn(100)
	}

}
