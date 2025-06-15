package main

import "fmt"

// Go is statically typed, so each variable has a type like:

// int, float64, bool, string  [ Discussed in this program ]

// Composite types like struct, array, slice, map, etc.

func main() {
	// There are Four types of variable declaration

	// Type 1: Explicit Type
	var myname string = "Pranav"
	//  Explicit declaration of varibale data type as string
	fmt.Println(myname)
	// ------------------------------------------------------------

	// Type 2: Inferred Type
	var myage = 23
	//  Here type is not delclared it is inferred by go as int
	fmt.Println(myage)
	// -------------------------------------------------------------

	// Type 3: Short Declaration
	Pi := 3.14
	// Here type is inferred as float and no need to say it is as variable
	// Note:- this works only inside functions
	fmt.Println(Pi)
	// ---------------------------------------------------------------

	// Type 4: Multiple varibles at once
	// var x, y, z int

	main2()
}

// ------------------------------------------------------------------------------
// When you declare a variable without initializing it, Go assigns a zero value based on the type:

// Type	                         Zero Value
// int	                           0
// float64	                       0.0
// string	                       ""
// bool	                           false
// pointers	                       nil

// ----------------------------------------------------------------------------------
// Scoping Rules

// Package-level variables are declared outside functions.
var packageLevel = "I am global"

// Function-level variables are declared inside a function.
func main2() {
	local := "I exist inside main"
	fmt.Println(packageLevel, local)
}
