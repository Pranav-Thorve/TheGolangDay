package main

import (
	"fmt"
	"time"
)

func checkWeekEnd() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's A WeekEnd")
	default:
		fmt.Println("It is a Week Day")
	}
}

func main() {
	checkWeekEnd()
	datatype := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its Interger")
		case string:
			fmt.Println("its a String")
		default:
			fmt.Println("pata nahi ji")
		}
	}

	datatype("pranav")
}
