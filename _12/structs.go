package main

import (
	"fmt"
)

// Struct Embeddings
type Work struct {
	title      string
	experience int
}

// Struct is a blueprint, you can create multple instance which are independent from each other
type Identity struct {
	name  string
	lname string
	age   int
	Work  // work struct emebedded into identity
}

// Constructors :- there is no constructior thing as such in go
func newIdentity(name string, lname string, age int) *Identity {
	myid := Identity{
		name:  name,
		lname: lname,
		age:   age,
	}

	return &myid
}

// Methods in struct
//
//	Receiver type
func (I *Identity) updateName(name string) {
	I.name = name
}

func main() {

	// mywork := Work{
	// 	title:      "DevOps/CloudOps Engineer",
	// 	experience: 2,
	// }
	myid := Identity{
		name:  "Pranav",
		lname: "Thorve",
		age:   23,
		// Work:  mywork,
		// inline
		Work: Work{
			"DevOps Engineer",
			2,
		},
	}
	// If any value is not assigned in a instance it gets the zero value

	fmt.Println(myid.lname)
	fmt.Println(myid)

	// can assign or change the value

	myid.age = 24
	fmt.Println(myid)

	myid.updateName("Pran")
	fmt.Println(myid)

	fmt.Println(newIdentity("a", "b", 3))

	// inline struct
	country := struct {
		name       string
		population int
	}{"india", 144}

	fmt.Println(country)

	// changing value ofembeededbstruct

	myid.Work.title = "CloudOps Engineer"

	fmt.Println(myid)

}
