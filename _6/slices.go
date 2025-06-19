package main

import (
	"fmt"
	"slices"
)

func Question1() {
	var q1 = []int{1, 2, 3, 4, 5}

	for i := range 5 {
		fmt.Println(q1[i])
	}
}

func Question2() {
	var names []string

	names = append(names, "Alice")
	names = append(names, "Bob")
	names = append(names, "Charlie")

	fmt.Println(names)
}

func Question3() {
	nums := []int{4, 9, 2, 7, 1, 10}

	slices.Sort(nums)
	fmt.Println(nums[len(nums)-1])
}

func Question4() {
	nums := []int{1, 2, 3, 4, 5}
	index := 2
	nums = slices.Delete(nums, index, index+1)
	fmt.Println(nums)
}

func Question5() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4, 5}

	for i := range 3 {
		if slices.Contains(b, a[i]) {
			index := slices.Index(b, a[i])
			b = slices.Delete(b, index, index+1)
		}
	}

	fmt.Println(slices.Concat(a, b))

}

func main() {
	// Declaring empty slice
	// uninitialized slice is nil
	var names []string
	fmt.Println(names == nil)

	// Declaring slices and initializing
	var nums = make([]int, 2, 3)
	//make(type,initial length, initial capacity)
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	// length and capacity can change
	// above length was 2 and capacity was 3
	// after appending elements
	// [0 0] initial slice
	nums = append(nums, 2)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 7)

	// length will be 7 and capacity will be 12
	// [0 0 2 4 5 6 7] slice after appending
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	// Copy function
	var num1 = make([]int, 0, 5)

	fmt.Println(num1)
	num1 = append(num1, 1)
	var num2 = make([]int, len(num1))
	copy(num2, num1)

	fmt.Println(num1, num2)

	// Slice Operator
	var numop = []int{1, 5, 7}

	fmt.Println(numop[0:2])
	//   [ from:to]
	//from start to specific index
	fmt.Println(numop[:1]) // prints elements from to index 1[excluded] , in this case only print 1
	// from specfic index to end
	fmt.Println(numop[1:]) // print elements from index 2[included] to last, in this case prints 5 and 7

	// Equal function
	var n1 = []int{1, 2, 3}
	var n2 = []int{1, 2, 4}
	var n3 = []int{1, 2, 3}

	fmt.Println(slices.Equal(n1, n2)) // print false as these slices differ at index 2 , 3!=4
	fmt.Println(slices.Equal(n1, n3)) // prints true , as all elements at index are equal

	// --------------------------------------------------------------------------------------------------------------

	// Q1. Create a slice of integers containing the numbers 1 to 5, and print each element using a for loop.
	Question1()

	// Q2. Start with an empty slice of strings. Append the names "Alice", "Bob", and "Charlie" to it. Print the final slice.
	Question2()

	//Q3 Given a slice of integers:
	// nums := []int{4, 9, 2, 7, 1, 10}
	// Write a function that returns the maximum value in the slice.
	Question3()

	// Q4 Write a function that removes an element at a given index from a slice of integers.
	//  Example input:
	// nums := []int{1, 2, 3, 4, 5}
	// index := 2
	Question4()

	//Q5 Write a function that takes two slices of integers and:

	// Merges them
	// Removes duplicate values
	// Returns the sorted result
	// Example Input:

	// a := []int{1, 2, 3}
	// b := []int{2, 3, 4, 5}
	Question5()
}
