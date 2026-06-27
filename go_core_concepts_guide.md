# Go Language Core Concepts — Beginner's Guide

A simple, example-driven reference covering Go concepts from pointers to goroutines.

---

## Table of Contents

1. [06 — Pointers](#06--pointers)
2. [07 — Arrays & Slices](#07--arrays--slices)
3. [08 — Maps](#08--maps)
4. [09 — Structs](#09--structs)
5. [10 — Methods](#10--methods)
6. [11 — Interfaces](#11--interfaces)
7. [12 — Packages](#12--packages)
8. [13 — Error Handling](#13--error-handling)
9. [14 — Goroutines](#14--goroutines)

---

## 06 — Pointers

### What is a Pointer?

A **pointer** stores the memory address of a variable instead of the value itself.  
Think of it like writing down someone's home address rather than bringing their whole house with you.

| Symbol | Meaning |
|--------|---------|
| `&x`   | "Give me the address of x" |
| `*p`   | "Give me the value at address p" |

### Example

```go
package main

import "fmt"

func main() {
    age := 25
    p := &age        // p holds the address of age

    fmt.Println(age)  // 25
    fmt.Println(p)    // 0xc000014080 (some memory address)
    fmt.Println(*p)   // 25  (value at that address)

    *p = 30           // change value through the pointer
    fmt.Println(age)  // 30  (age changed too!)
}
```

### Why Use Pointers?

```go
// WITHOUT pointer — original value NOT changed
func addTen(n int) {
    n = n + 10
}

// WITH pointer — original value IS changed
func addTenPtr(n *int) {
    *n = *n + 10
}

func main() {
    x := 5
    addTen(x)
    fmt.Println(x)    // 5 (unchanged)

    addTenPtr(&x)
    fmt.Println(x)    // 15 (changed!)
}
```

> **Key Rule:** Use pointers when you want a function to modify the original variable, or when passing large structs to avoid copying them.

---

## 07 — Arrays & Slices

### Arrays — Fixed Size

An array has a **fixed length** that cannot change after declaration.

```go
package main

import "fmt"

func main() {
    // Declare an array of 3 integers
    var scores [3]int
    scores[0] = 10
    scores[1] = 20
    scores[2] = 30

    // Short declaration
    names := [3]string{"Alice", "Bob", "Charlie"}

    fmt.Println(scores)         // [10 20 30]
    fmt.Println(names)          // [Alice Bob Charlie]
    fmt.Println(len(names))     // 3
}
```

### Slices — Dynamic Size ⭐

A **slice** is like an array but it can grow or shrink. Most Go code uses slices, not arrays.

```go
package main

import "fmt"

func main() {
    // Create a slice
    fruits := []string{"apple", "banana", "cherry"}

    // Add elements with append
    fruits = append(fruits, "mango")
    fmt.Println(fruits)         // [apple banana cherry mango]
    fmt.Println(len(fruits))    // 4 (current length)
    fmt.Println(cap(fruits))    // 6 (underlying capacity)

    // Slice a slice  [start : end]  (end is exclusive)
    part := fruits[1:3]
    fmt.Println(part)           // [banana cherry]

    // make() — create slice with length and capacity
    numbers := make([]int, 3, 5)
    fmt.Println(numbers)        // [0 0 0]
}
```

### Looping Over a Slice

```go
colors := []string{"red", "green", "blue"}

for index, value := range colors {
    fmt.Printf("Index %d: %s\n", index, value)
}
// Index 0: red
// Index 1: green
// Index 2: blue
```

> **Quick Comparison:**
> - **Array** → `[3]int{1, 2, 3}` — fixed, rarely used directly
> - **Slice** → `[]int{1, 2, 3}` — flexible, used everywhere

---

## 08 — Maps

### What is a Map?

A **map** stores key-value pairs — like a dictionary or lookup table.  
Format: `map[KeyType]ValueType`

```go
package main

import "fmt"

func main() {
    // Create a map
    person := map[string]int{
        "Alice": 30,
        "Bob":   25,
        "Carol": 28,
    }

    // Access a value
    fmt.Println(person["Alice"])  // 30

    // Add a new key
    person["Dave"] = 22

    // Update existing key
    person["Bob"] = 26

    // Delete a key
    delete(person, "Carol")

    fmt.Println(person)
    // map[Alice:30 Bob:26 Dave:22]
}
```

### Check if a Key Exists

```go
age, exists := person["Alice"]
if exists {
    fmt.Println("Alice's age:", age)
} else {
    fmt.Println("Alice not found")
}
```

### Loop Over a Map

```go
for name, age := range person {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

### Create with make()

```go
// Using make
scores := make(map[string]int)
scores["math"] = 95
scores["english"] = 88
```

> **Note:** Maps are unordered — the iteration order is random every time.

---

## 09 — Structs

### What is a Struct?

A **struct** groups related data together under one name.  
Think of it as a custom data type — like a blueprint for an object.

```go
package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Create a struct instance
    p1 := Person{
        Name: "Alice",
        Age:  30,
        City: "Mumbai",
    }

    // Access fields with dot notation
    fmt.Println(p1.Name)   // Alice
    fmt.Println(p1.Age)    // 30

    // Modify a field
    p1.Age = 31
    fmt.Println(p1)        // {Alice 31 Mumbai}
}
```

### Struct with Pointer

```go
p2 := &Person{Name: "Bob", Age: 25, City: "Delhi"}
p2.Age = 26   // Go automatically dereferences — no need for (*p2).Age
fmt.Println(p2.Age)  // 26
```

### Nested Structs

```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Name    string
    Address Address   // nested struct
}

func main() {
    e := Employee{
        Name: "Ravi",
        Address: Address{Street: "MG Road", City: "Bangalore"},
    }
    fmt.Println(e.Address.City)  // Bangalore
}
```

> **Struct vs Map:** Use a **struct** when you know the fields in advance and want type safety. Use a **map** for dynamic or unknown keys.

---

## 10 — Methods

### What is a Method?

A **method** is a function attached to a specific type (usually a struct).  
It lets you define behavior for your custom types.

```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// Method on Rectangle — note (r Rectangle) before function name
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    fmt.Println(rect.Area())       // 15
    fmt.Println(rect.Perimeter())  // 16
}
```

### Pointer Receiver — Modify the Struct

Use a **pointer receiver** `*Type` when the method needs to change the struct's data.

```go
type Counter struct {
    Count int
}

// Value receiver — does NOT change original
func (c Counter) ShowCount() {
    fmt.Println("Count:", c.Count)
}

// Pointer receiver — DOES change original
func (c *Counter) Increment() {
    c.Count++
}

func main() {
    c := Counter{Count: 0}
    c.Increment()
    c.Increment()
    c.ShowCount()  // Count: 2
}
```

> **Rule of thumb:**
> - Use **value receiver** `(t Type)` → for read-only methods
> - Use **pointer receiver** `(t *Type)` → when the method modifies the struct

---

## 11 — Interfaces

### What is an Interface?

An **interface** defines a set of method signatures that a type must implement.  
It enables polymorphism — different types can be used interchangeably.

```go
package main

import (
    "fmt"
    "math"
)

// Define an interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Circle implements Shape
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Rectangle implements Shape
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Function that accepts any Shape
func printInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 3}

    printInfo(c)   // Area: 78.54, Perimeter: 31.42
    printInfo(r)   // Area: 12.00, Perimeter: 14.00
}
```

### Empty Interface

`interface{}` (or `any` in Go 1.18+) accepts any type.

```go
func printAnything(v interface{}) {
    fmt.Println(v)
}

printAnything(42)
printAnything("hello")
printAnything(true)
```

> **Key Point:** In Go, interfaces are implemented **implicitly** — you don't need to write `implements Shape`. If a type has all the required methods, it automatically satisfies the interface.

---

## 12 — Packages

### What is a Package?

A **package** is a collection of Go files in the same directory that share the same `package` name.  
Packages let you organize and reuse code.

### Using Built-in Packages

```go
package main

import (
    "fmt"      // formatted I/O
    "math"     // math functions
    "strings"  // string utilities
    "strconv"  // string conversions
)

func main() {
    // fmt
    fmt.Printf("Hello, %s! You are %d years old.\n", "Alice", 30)

    // math
    fmt.Println(math.Sqrt(16))     // 4
    fmt.Println(math.Abs(-7.5))    // 7.5

    // strings
    fmt.Println(strings.ToUpper("hello"))          // HELLO
    fmt.Println(strings.Contains("golang", "go"))  // true
    fmt.Println(strings.Replace("aabbcc", "b", "x", -1))  // aaxxcc

    // strconv
    num, _ := strconv.Atoi("42")   // string to int
    fmt.Println(num + 1)           // 43
    fmt.Println(strconv.Itoa(100)) // "100"
}
```

### Creating Your Own Package

**File: `greet/greet.go`**
```go
package greet

import "fmt"

// Exported function — starts with Capital letter
func Hello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// unexported function — starts with lowercase (only usable inside package)
func secret() {
    fmt.Println("this is private")
}
```

**File: `main.go`**
```go
package main

import "yourmodule/greet"

func main() {
    greet.Hello("Alice")  // Hello, Alice!
}
```

> **Visibility Rule:**
> - **Uppercase** first letter → exported (public)
> - **Lowercase** first letter → unexported (private to the package)

---

## 13 — Error Handling

### How Go Handles Errors

Go does **not** use exceptions. Instead, functions return an `error` value as the last return value.  
You check it after every call.

```go
package main

import (
    "errors"
    "fmt"
)

// Function that returns a value AND an error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil   // nil means no error
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)  // Result: 5
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)  // Error: cannot divide by zero
    }
}
```

### fmt.Errorf — Formatted Error Messages

```go
func getUser(id int) (string, error) {
    if id <= 0 {
        return "", fmt.Errorf("invalid user id: %d", id)
    }
    return "Alice", nil
}
```

### Custom Error Types

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on '%s': %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{Field: "age", Message: "must be non-negative"}
    }
    return nil
}

func main() {
    err := validateAge(-5)
    if err != nil {
        fmt.Println(err)
        // validation failed on 'age': must be non-negative
    }
}
```

### panic and recover

Use `panic` for truly unrecoverable situations, and `recover` to catch panics gracefully.

```go
func safeDiv(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
            result = -1
        }
    }()
    return a / b  // panics if b == 0
}

func main() {
    fmt.Println(safeDiv(10, 2))   // 5
    fmt.Println(safeDiv(10, 0))   // Recovered from panic: ... then -1
}
```

> **Best Practice:** Prefer returning errors over panicking. Reserve `panic` for truly unexpected situations (like programmer bugs, not user input errors).

---

## 14 — Goroutines

### What is a Goroutine?

A **goroutine** is a lightweight thread managed by the Go runtime.  
You can run thousands of goroutines concurrently with very little memory overhead.

Start a goroutine simply by adding `go` before a function call.

```go
package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    fmt.Printf("Hello from %s!\n", name)
}

func main() {
    go sayHello("Alice")   // runs concurrently
    go sayHello("Bob")     // runs concurrently
    go sayHello("Charlie") // runs concurrently

    time.Sleep(1 * time.Second)  // wait for goroutines to finish
    fmt.Println("Done!")
}
```

### Channels — Communication Between Goroutines

A **channel** is a pipe that goroutines use to safely send and receive values.

```go
package main

import "fmt"

func sum(nums []int, ch chan int) {
    total := 0
    for _, n := range nums {
        total += n
    }
    ch <- total  // send result into channel
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    ch := make(chan int)

    // Split work between two goroutines
    go sum(numbers[:5], ch)   // sums 1+2+3+4+5
    go sum(numbers[5:], ch)   // sums 6+7+8+9+10

    part1 := <-ch  // receive from channel
    part2 := <-ch  // receive from channel

    fmt.Println("Total:", part1+part2)  // Total: 55
}
```

### WaitGroup — Wait for All Goroutines

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()   // signals this goroutine is done
    fmt.Printf("Worker %d starting\n", id)
    // ... do work ...
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)           // tell WaitGroup to expect one more goroutine
        go worker(i, &wg)
    }

    wg.Wait()  // block until all goroutines call Done()
    fmt.Println("All workers finished!")
}
```

### Select — Multiplex Channels

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "one"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "two"
    }()

    // select waits for whichever channel is ready first
    for i := 0; i < 2; i++ {
        select {
        case msg := <-ch1:
            fmt.Println("Received:", msg)
        case msg := <-ch2:
            fmt.Println("Received:", msg)
        }
    }
}
```

> **Goroutine Tips:**
> - Use `sync.WaitGroup` when you don't need data back from goroutines
> - Use **channels** when goroutines need to share data
> - Use `select` when listening on multiple channels at once

---

## Quick Concept Summary

| Concept | One-Line Description |
|---------|---------------------|
| **Pointers** | Store memory addresses; use `&` to get address, `*` to dereference |
| **Arrays** | Fixed-size sequences — `[3]int{1, 2, 3}` |
| **Slices** | Dynamic arrays — grow with `append()`, slice with `[start:end]` |
| **Maps** | Key-value store — `map[string]int{"a": 1}` |
| **Structs** | Group related fields into a custom type |
| **Methods** | Functions tied to a type — use pointer receiver to modify |
| **Interfaces** | Contract of methods — implemented implicitly |
| **Packages** | Organize code; uppercase = exported, lowercase = private |
| **Error Handling** | Return `error` as last value; check `if err != nil` |
| **Goroutines** | Lightweight concurrent functions launched with `go` keyword |

---