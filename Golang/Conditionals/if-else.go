/*
   If statement conditions

   If-else conditional statement is used in Golang when a situation leads to two conditions and one of them should hold true.
   Golang relies on indentation to define scope in the code
*/

/*
    Syntax:

    if condition {
	    code1
    } else if {
	    code2
    } else {
	    code3
    }
*/

package main

import "fmt"

func main() {
	// Example
	x := 69
	y := 19
	if x > y {
		fmt.Println("b is greater than a")
	} else if x > y {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("Both are Equal")
	}

	/*
	   &&  , || keywords

	   Used to combine conditional statements

	*/

	// && Example

	a := 200
	b := 33
	c := 500
	if a > b && c > a {
		fmt.Println("And is used when Both statements are true")
	}

	// || Example

	a = 200
	b = 33
	c = 500
	if a > b || a > c {
		fmt.Println("or is used when At least one of the conditions is True")
	}

	/*
	   Nested IF

	   if statements inside if statements

	*/

	// Example

	x = 41
	if x > 10 {
		fmt.Printf("Above ten, ")
		if x > 20 {
			fmt.Println("and also above 20!")
		} else {
			fmt.Println("but not above 20.")
		}
	}
}

// Output
// b is greater than a
// And is used when Both statements are true
// or is used when At least one of the conditions is True
// Above ten, and also above 20!
