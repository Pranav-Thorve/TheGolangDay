package main

import "fmt"

// A constant in Go is a named value that does not change

func constants() {
	// They can declared explicitel or can be inferred by go itself
	// Type1: untyped constants
	const Pi = 3.14

	// Type2: typed constants
	const Greeting string = "Hello, world"

	// Type3: Multiple.constants
	const (
		PI = 3.14
		E  = 2.71
		G  = 9.8
	)

	// Constants can be the result of compile-time expressions:
	const msg = "Hi, " + "go!"

	// IOTA
	// iota is a special keyword used to create incremental constants (usually in enums).
	const (
		A = iota // 0
		B        // 1
		C        // 2
	)

	fmt.Println(C)

}
