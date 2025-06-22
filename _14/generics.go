package main

import "fmt"

//  Generics in go defines any type of value or defined type of
// values can be accpepted for functions or structs

// [T any] this can be
// [T any] , [T string | any ] , [T comparable]
// [T any] => any type
// [T string | any ] => only string and int atr allwed
// [T comparable] => all comparable type (booleans, numbers,strings,pointers,channels,arrays of comparable types, structs whose fields are comparable tyoes)
// can be suppplied mutple generics [T any, V string]
func printanything[T any](item T) {
	fmt.Println(item)
}

func main() {
	printanything("name")
	printanything(2)
	printanything(true)
}
