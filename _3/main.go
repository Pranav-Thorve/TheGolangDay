package main

import (
	"fmt"
)

// Q1. Write a Go program that uses a for loop to calculate the sum of the first 10 natural numbers and print the result.

func sumOfFirstNNatualNumbers() {
	// Declared a variable which stores sum of numbers
	var sum int

	// one of the way to write for loop with the help of range
	for i := range 11 {
		sum += i
		i = i + 1
	}

	println(sum)
}

// Q2. Write a program that loops through numbers from 1 to 20, and for each number:

// Print "Even" if the number is even
// Print "Odd" if the number is odd

func isItEvenOrOdd() {

	var i int = 1

	for i <= 20 {
		if i%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
		i = i + 1
	}
}

// Q3. Write a program that prints the multiplication table of a number (say, 5) using a for loop.

func multiplicationTable(number int) {
	for i := range 11 {
		if i == 0 {
			continue
		}
		fmt.Println(number, "x", i, "=", number*i)
	}
}

// Q4. Write a program using nested for loops and if-else to print all prime numbers between 1 and 50.

func suppDividefunc(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func isPrime() {
	for i := 1; i <= 50; i++ {
		if suppDividefunc(i) {
			fmt.Println(i, "is Prime")
		}
	}
}

// Q5. Write a program that loops from 1 to 100 and:

// Prints "Fizz" if the number is divisible by 3
// Prints "Buzz" if divisible by 5
// Prints "FizzBuzz" if divisible by both 3 and 5
// Otherwise, print the number

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {

	// sumOfFirstNNatualNumbers()
	// isItEvenOrOdd()
	// multiplicationTable(5)
	// isPrime()
	fizzbuzz()
}
