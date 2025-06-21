package main

import "fmt"

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp

}

func makeAdder(num int) func(int) int {

	value := func(gnum int) int {
		return num + gnum
	}

	return value
}

func appendToSlice(s *[]int, num int) {
	*s = append(*s, num)
}

func newCounter() func() int {
	num := 0

	value := func() int {
		num = num + 1
		return num
	}

	return value

}

func applyFunc(n *int, f func(int) int) {
	*n = f(*n)
}

func main() {

	// Q1. Write a function that swaps the values of two integers using pointers.

	// func swap(a *int, b *int)
	x, y := 5, 10
	fmt.Println(x, y)
	Swap(&x, &y)
	fmt.Println(x, y)

	// 	Q2. Write a function makeAdder that returns a closure which adds a given number to its input.

	add5 := makeAdder(5)
	fmt.Println(add5(3)) // Should print 8

	//Q3. Write a function that appends a number to a slice using a pointer.
	var nums = []int{1, 2, 3}
	fmt.Println(nums)
	appendToSlice(&nums, 5)
	fmt.Println(nums)

	// Q4. Write a function newCounter that returns a closure. Each time you call the returned function, it increments and returns an internal counter.

	count := newCounter()
	fmt.Println(count()) // 1
	fmt.Println(count()) // 2
	fmt.Println(count()) // 3

	// Q5. Write a function applyFunc that takes:

	// A pointer to an integer
	// A function that takes and returns an integer
	// It applies the function to the value via pointer.
	num := 5
	applyFunc(&num, func(x int) int { return x * 2 })
	fmt.Println(num)

}
