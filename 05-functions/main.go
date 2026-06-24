package main

import "fmt"

// 1 :: Simple Function
func greet() {
	fmt.Println("Hello World!")
}

// 2 :: Parameterized Function
func greetWithName(name string) {
	fmt.Println("Hello,", name)
}

// 3 :: Parameterized Function with Return Type
func greetToName(name string) string {
	return "Hello, " + name
}

// 4 :: Multiple Return Values
func multiValueReturn() (string, int) {
	return "Hi", 200
}

// 5 :: Named Function Assigned to Variable
func multiply(a int, b int) int {
	return a * b
}

// 6 :: Variadic Function
func sumAll(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// 7 :: Function Returning Function
func f1() func() {
	return func() {
		fmt.Println("Hello from returned function!")
	}
}

// 8 :: Function as Parameter
func operate(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

// 9 :: Closure
func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// 10 :: Recursive Function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

// 11 :: Method
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 12 :: Pointer Receiver Method
type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

// Bonus :: Named Return Value
func divide(a int, b int) (result int) {
	result = a / b
	return
}

// Bonus :: Defer
func testDefer() {
	defer fmt.Println("Deferred Statement")

	fmt.Println("Normal Statement")
}

func main() {

	fmt.Println("######## Function Practice ########")

	// 1
	greet()

	// 2
	greetWithName("Harry")

	// 3
	fmt.Println(greetToName("Alex"))

	// 4
	msg, code := multiValueReturn()
	fmt.Println(msg, code)

	// Anonymous Function
	sum := func(a int, b int) int {
		return a + b
	}

	fmt.Println(sum(10, 15))

	// 5
	mul := multiply
	fmt.Println(mul(10, 20))

	// 6
	fmt.Println(sumAll(1, 2))
	fmt.Println(sumAll(1, 2, 3, 4, 5))

	// 7
	xyz := f1()
	xyz()

	// 8
	fmt.Println(operate(10, 30, sum))

	// 9 :: Closure
	c := counter()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	// 10 :: Recursion
	fmt.Println("Factorial:", factorial(5))

	// 11 :: Method
	rect := Rectangle{
		Width:  10,
		Height: 20,
	}

	fmt.Println("Area:", rect.Area())

	// 12 :: Pointer Receiver Method
	ctr := Counter{}

	ctr.Increment()
	ctr.Increment()

	fmt.Println("Counter:", ctr.Value)

	// Bonus :: IIFE
	result := func(a int, b int) int {
		return a + b
	}(100, 200)

	fmt.Println("IIFE Result:", result)

	// Bonus :: Named Return
	fmt.Println("Division:", divide(20, 5))

	// Bonus :: Defer
	testDefer()
}