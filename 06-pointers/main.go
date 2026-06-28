package main

import "fmt"

// Call by Value
func callByValue(p int) {
	p = 100
}

// Call by Reference (using pointer)
func callByReference(p *int) {
	*p = 100
}

func main() {

	fmt.Println("========== Pointer Basics ==========")

	x := 10

	var p *int
	p = &x

	fmt.Println("Value of x      :", x)
	fmt.Println("Address of x    :", &x)
	fmt.Println("Pointer p       :", p)
	fmt.Println("Value using *p  :", *p)

	fmt.Println()

	fmt.Println("Changing value using pointer...")
	*p = 50

	fmt.Println("Updated x       :", x)

	fmt.Println()

	fmt.Println("========== Call by Value ==========")

	callByValue(x)

	fmt.Println("Value of x :", x)

	fmt.Println()

	fmt.Println("========== Call by Reference ==========")

	callByReference(&x)

	fmt.Println("Value of x :", x)

	fmt.Println()

	fmt.Println("========== Nil Pointer ==========")

	var ptr *int

	fmt.Println(ptr)

	fmt.Println()

	fmt.Println("========== Using new() ==========")

	num := new(int)

	fmt.Println(*num)

	*num = 500

	fmt.Println(*num)
}