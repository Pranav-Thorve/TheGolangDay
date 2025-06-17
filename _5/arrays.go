package main

import "fmt"

// Arrays are fixed size
// use when you know the exact size, it is better for memory optimization

func main() {
	var numbers [2]int
	// Zero array int - 0, bool- false, string-""

	// assign a value
	numbers[1] = 2
	fmt.Println(numbers)

	// assigning values at declaration

	names := [2]string{"pranav", "radhey"}

	fmt.Println(names[0])
}
