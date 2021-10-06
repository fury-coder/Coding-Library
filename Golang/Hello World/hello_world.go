// Single line comment
/* Multi-
 line comment */

// A package clause starts every source file.
// Main is a special name declaring an executable rather than a library.
package main

// Import declaration declares library packages referenced in this file.
import "fmt" // A package in the Go standard library.

// A function definition. Main is special. It is the entry point for the
// executable program. Love it or hate it, Go uses brace brackets.
func main() {
    // Println outputs a line to stdout.
    // It comes from the package fmt.
	fmt.Println("Hello World!")
}



//Output
//Hello World!
