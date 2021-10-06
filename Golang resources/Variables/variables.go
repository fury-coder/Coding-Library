/*
Variables:

A typical program uses various values that may change during its execution. For Example, a program that performs some operations on the values entered by the user. The values entered by one user may differ from those entered by another user. Hence this makes it necessary to use variables as another user may not use the same values. When a user enters a new value that will be used in the process of operation, can store temporarily in the Random Access Memory of the computer and these values in this part of memory vary throughout the execution and hence another term for this came which is known as Variables. So basically, a Variable is a placeholder of the information which can be changed at runtime. And variables allow to Retrieve and Manipulate the stored information.
*/

/* Syntax
var variable_name type = expression
*/

// Short hand Syntax
// variable_name := explression

package main

import "fmt"

func main() {
    // Variable declared and initialized without expression
	var variable1 = "hello"
	fmt.Printf("%T\n", variable1)

    // Variable declared and initialized with expression
	var variable2 int = 10
	fmt.Printf("%T\n", variable2)

    // Variable declared and initialized without expression
	var variable3 float32
	variable3 = 23.0
	fmt.Printf("%T\n", variable3)

	// Multiple variables of the same type are declared and initialized in the single line	
    var variable4, variable5, variable6 int = 2, 454, 67
    fmt.Printf("%T %T %T\n", variable4, variable5, variable6) 

    // Multiple variables of the same type are declared and initialized in the single line
    var variable7, variable8, variable9 = 2, "hello", 67.56
    fmt.Printf("%T %T %T\n", variable7, variable8, variable9)

    // Variable declared using the short-hand way
    variable10 := 69
    fmt.Printf("%T\n", variable10)
    
}

// Output
// string
// int
// float32
// int int int
// int string float64
// int
