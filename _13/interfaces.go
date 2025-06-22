package main

import (
	"fmt"
)

// INterfaces are like contratcs
//  in go Interfaces are implemented implicielty,
// we dont have to use keyword like implments on strictt
// but method signature should be same

type Transporter interface {
	ride()
}

type Transport struct {
	vehicle Transporter
}

func (T Transport) takeRide() {
	T.vehicle.ride()
}

type Bus struct{}

func (b Bus) ride() {
	fmt.Println("Riding by bus")
}

type Train struct{}

func (t Train) ride() {
	fmt.Println("Riding by Train")
}

func main() {

	trainride := Train{}
	myride := Transport{
		vehicle: trainride,
	}

	myride.takeRide()
}
