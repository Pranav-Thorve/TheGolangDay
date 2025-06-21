package main

import "fmt"

// Thing 1 :- Normal Function Declaration
func Normal() {
	fmt.Println("Normal")
}

// Thing 2 :- Accepting parameters

func Acceptor(a, b int) {
	fmt.Println(a + b)
}

// Thing 3 :- Returning value
func ReturnValue(a, b int) int {
	return a + b
}

// Thing 4 :- Returning Multiple values

func MultiReturn() (int, string) {
	return 2, "name"
}

// Thing 5 :- Accepting any type of value

func AnyParameters(num interface{}) {
	fmt.Println(num)
}

// Thing 6 :- Accepting N Parameters, It is called variadic function
func NParameters(nums ...int) {

}

// Thing 7:- Returning a function
func ReturnFunction() func(int) int {
	return func(a int) int {
		return 3
	}
}

// Thing 8 :- Accepting function

func FunctionAcceptor(fn func(a int) int) {
	fn(1)
}

func main() {
	Normal()
	Acceptor(2, 3)
	ReturnValue(2, 3)
	val1, val2 := MultiReturn()
	fmt.Println(val1, val2)

	AnyParameters(2)
	AnyParameters("name")

	mslice := []int{12, 3, 4, 4, 5}
	NParameters(mslice...)

	ReturnFunction()

	fn := func(a int) int {
		return a
	}
	FunctionAcceptor(fn)
}
