package main

import (
	"fmt"
	"maps"
)

func Question1() {
	mp := map[string]int{"apple": 5, "banana": 3, "orange": 7}

	fmt.Println(mp["banana"])
}

func Question2() {
	mp := make(map[string]int)

	mp["pen"] = 10
	mp["pencil"] = 15

	// update pen to 12
	mp["pen"] = 12
	delete(mp, "pencil")

	fmt.Println(mp)
}

func Question3(myslice []string) map[string]int {
	// []string{"go", "java", "go", "python", "go", "java"}
	rmap := make(map[string]int)

	for i := range len(myslice) {
		rmap[myslice[i]]++
	}
	return rmap
}

func Question4(key string) {
	users := map[string]string{"alice": "123", "bob": "qwerty"}
	_, ok := users[key]

	if ok {
		fmt.Println("user found")
	} else {
		fmt.Println("user not found")
	}
}

func Question5() {
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	newmap := make(map[int]string)

	for key, value := range original {
		newmap[value] = key
	}

	fmt.Println(newmap)
}

func main() {

	// Declaring a Map
	mymap := make(map[int]int)
	//map[key type]value type

	// setting an element
	mymap[2] = 4

	fmt.Println(mymap)

	// Getting an elemnt
	fmt.Println(mymap[2])
	// if key does not exist then the map return zeroed values -> int-0, string -"", bool-false
	fmt.Println(mymap[7])
	mymap[3] = 6
	// Delete and clear function
	fmt.Println(mymap[3])
	delete(mymap, 3)
	//(map,key)
	fmt.Println(mymap[3])
	clear(mymap)
	fmt.Println(mymap)

	// Takeaway;- Delete functiona deletes a specific key value, clear function empty the map

	// Declaring and initializing map at a time

	mp := map[string]int{"age": 23, "year": 2001}

	fmt.Println(mp)

	// Comma Ok idio,

	_, ok := mp["age"]
	//value,ok[true or falsr] :=
	v, ok := mp["year"]
	fmt.Println(v)
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	// Equal function
	mp1 := map[string]int{"age": 23, "year": 2001}
	mp2 := map[string]int{"age": 23, "year": 2001}

	fmt.Println(maps.Equal(mp1, mp2))

	//Q1 Create a map with the following key-value pairs:

	// "apple": 5
	// "banana": 3
	// "orange": 7
	// Then, print the quantity of "banana".
	Question1()

	//Q2. Create an empty map of type map[string]int.

	// Add keys "pen" (10) and "pencil" (15)
	// Update "pen" to 12
	// Delete "pencil"
	// Print the final map
	Question2()

	// Write a function that takes a slice of strings like:

	// []string{"go", "java", "go", "python", "go", "java"}
	// and returns a map with the count of each string.
	fmt.Println(Question3([]string{"go", "java", "go", "python", "go", "java"}))

	// Q4. Given a map of usernames and passwords:

	// users := map[string]string{"alice": "123", "bob": "qwerty"}
	// Write a function that:

	// Takes a username as input
	// Checks if it exists in the map using the comma-ok idiom
	// Prints "User found" or "User not found"
	Question4("alice")

	// Q5. Given a map:

	// original := map[string]int{"a": 1, "b": 2, "c": 3}
	// Create a new map that reverses keys and values:

	// map[int]string{1: "a", 2: "b", 3: "c"}
	Question5()
}
