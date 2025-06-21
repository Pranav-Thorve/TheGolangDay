package main

import "fmt"

// Baasics:-
// let say you have a variale with a value and somewhere in the function you declare a same
//  variable but now you want to change the value of the variable which
//  was declared before , so basically pointers are the like two variables pointing to same memory
//  locations to store and get value

func somewhere(num *int) {
	*num = 5
	fmt.Println("from somewhere", *num)
}

func main() {
	num := 2

	somewhere(&num)

	fmt.Println("from main", num)
}
