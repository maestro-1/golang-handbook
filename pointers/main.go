package main

import "fmt"

/*
The module intro deals with pointers

Deref is how you change/update the value of non pointer values in functions, etc without returning and reassigning
*/

func updateName(x string)  {
	x = "wedge"
}

// * means accepts pointer of a value
func updateNameDeref(x *string)  {
	// This update the value of the pointer without returning and reassigning when called
	*x = "wedge"
}

func introToPointer()  {
	name := "tifa"
	updateName(name)

	m := &name
	fmt.Println("memory address of name is: ",m)

	// * in front of a pointer gets the value of said pointer. Called a deref value
	fmt.Println("value at memory address: ",*m)
}

func introToDeref()  {
	name := "tifa"
	
	fmt.Println("memory address of name is: ",name)

	m := &name
	updateNameDeref(m)
	fmt.Println("memory address of name is: ",name)
}

func main()  {
	introToPointer()
	introToDeref()
}
