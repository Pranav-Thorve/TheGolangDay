package main

import "fmt"

// basics:-
// basically closure is a function defined inside a funtion amd
//  can access the varibles defined in outer function without
//  changing its value
//  What happens is it stores the copy of that variable and function inside separate memory localtion for a function call

func outer() func(int) int {
	age := 23

	value := func(num int) int {
		age = age - num
		return age
	}

	return value
}

func main() {
	firstcall := outer()
	secondcall := outer()

	fmt.Println(firstcall(3))
	//  when i gave i first call inner function used the value age and decresed
	//  value by 3 so now when i ll give the secondcall it should decrese the value by 5 more which will result in 15
	//  but it does not it give me 18 how??
	//  whne i give a first call the closure of a variable and a inner function is stored in another memory localtion
	//   and and when i gave a second call again stores a variable and a function in aother memory localtion with the same value as 23
	fmt.Println(secondcall(5))
}
