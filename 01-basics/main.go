package main

import "fmt"

func main() {

	// ==============================
	// Printing to the Console
	// ==============================

	fmt.Println("Hello World!")
	fmt.Printf("My name is %v\n", "Bruce Wayne")
	fmt.Print("Nice to meet you. ", "How are you?\n")

	// Format Specifiers
	// %v -> Prints the value
	// %T -> Prints the data type

	// ==============================
	// Variables
	// ==============================

	var name = "Jon Snow"              // Type inferred
	name2 := "Harry Potter"            // Short variable declaration
	var name3 string = "Pavan Kumar"   // Explicit type

	// Print multiple variables
	fmt.Println(name, name2, name3)

	// Print data types
	fmt.Printf("Type of name  : %T\n", name)
	fmt.Printf("Type of name2 : %T\n", name2)
	fmt.Printf("Type of name3 : %T\n", name3)

	// NOTE:
	// Every declared variable and imported package must be used.
	// Otherwise, the compiler throws an error.

	// ==============================
	// Constants
	// ==============================

	const Pi = 3.14

	fmt.Println("Value of Pi:", Pi)

	// NOTE:
	// Unused constants do NOT cause a compiler error.

	// ==============================
	// Zero Values
	// ==============================

	var fName string

	// Zero value of string is an empty string ("")
	fmt.Println("Zero value of fName:", fName)

	fName = "Ravi"

	fmt.Println("Updated fName:", fName)

	// ==============================
	// Arrays
	// ==============================

	// Arrays have a fixed size.

	var fruits = [5]string{}

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[4] = "Orange"

	// Prints the complete array (including empty elements)
	fmt.Println("Fruits:", fruits)

	// Print array type
	fmt.Printf("Type of fruits: %T\n", fruits)

	// ==============================
	// Built-in len() Function
	// ==============================

	fmt.Println("Array Length :", len(fruits))
	fmt.Println("String Length:", len(fName))

	// NOTE:
	// Arrays have a fixed size.
	// Slices are dynamic and are used more often in Go.

	// ==============================
	// Different Ways to Declare Arrays
	// ==============================

	// 1. Declare an empty array and assign values later.
	var arr [5]string
	arr[0] = "Hello"
	fmt.Println(arr)

	// 2. Initialize an empty array.
	var arr2 = [2]string{}
	arr2[0] = "Hi"
	fmt.Println(arr2)

	// 3. Initialize with values.
	var arr3 = [2]string{"Bye"}
	fmt.Println(arr3)

	// 4. Short variable declaration.
	arr4 := [3]string{"XYZ"}
	fmt.Println(arr4)

	// ==============================
	// Slices
	// ==============================

	// Slices are dynamic arrays.
	// append() returns a new slice, so always assign it back.
	// A slice internally stores:
	// - Pointer to the underlying array
	// - Length
	// - Capacity

	var s1 = []string{}

	s1 = append(s1, "Hello World")
	s1 = append(s1, "Bye Bye!")

	fmt.Println(s1)

	fmt.Printf("Length   : %d\n", len(s1))
	fmt.Printf("Capacity : %d\n", cap(s1))
	fmt.Printf("Type     : %T\n", s1)

	// NOTE:
	// If append() exceeds the slice capacity,
	// Go allocates a new underlying array and
	// returns a new slice pointing to that array.

	// ==============================
	// Compiler Errors
	// ==============================

	// Uncommenting either of these will produce a compile error
	// because local variables must be used.

	// var s2 []string
	// var a = [2]string{}
}