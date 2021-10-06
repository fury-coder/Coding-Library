/*
Data types specify the type of data that a valid Go variable can hold. In Go language, the type is divided into four categories which are as follows:

    The Basic data types are as follows:
    Integers: int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint, rune, byte, uintpt
    Float: float32, float64
    complex: complex64, complex128
    Booleans: true, false
    Strings
*/
package main

import "fmt"

func main() {
	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	// Using floating-point number
	a := 20.45
	fmt.Printf("Result is: %f", a)

	// Using complex numbers
	var a1 complex128 = complex(6, 2)
	fmt.Println(a1)

	// using Booleans
	str1 := 1
	str2 := 2
	str3 := 1
	fmt.Println(str1 == str2) // false
	fmt.Println(str1 == str3) // true

	str := "Hello World"
	// Display the length of the string
	fmt.Printf("Length of the string is:%d",
		len(str))
	// Display the string
	fmt.Printf("\nString is: %s", str)
}

// Output
// 225 222
// -32767 32765
// Result is: 20.450000(6+2i)
// false
// true
// Length of the string is:11
// String is: Hello World
