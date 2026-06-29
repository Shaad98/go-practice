# Go Arrays & Slices

This document covers the fundamentals of **Arrays** and **Slices** in Go, including declaration, initialization, traversal, copying, slicing, insertion, deletion, capacity management, and common operations used in real-world Go applications.

---

# Arrays

## What is an Array?

An **array** is a fixed-size collection of elements of the same type.

* Size is decided during declaration.
* Memory is allocated immediately.
* All elements are stored contiguously in memory.
* Arrays are value types.

### Declaration

```go
var numbers [5]int
```

Creates an array capable of storing **5 integers**.

Default values:

```
[0 0 0 0 0]
```

---

## Initialization

```go
arr := [5]int{10, 20, 30, 40, 50}
```

Creates an array and initializes all elements.

---

## Accessing Elements

Access an element using its index.

```go
arr[0]
```

Output

```
10
```

Index starts from **0**.

---

## Updating Elements

```go
arr[2] = 100
```

Updates the value at index `2`.

---

## Length

```go
len(arr)
```

Returns the total number of elements.

Output

```
5
```

---

## Iterating Over Arrays

### Traditional for loop

```go
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

Useful when index manipulation is required.

---

### range loop

```go
for index, value := range arr {
    fmt.Println(index, value)
}
```

`range` returns:

* Index
* Value

If index isn't needed:

```go
for _, value := range arr {
    fmt.Println(value)
}
```

---

## Copying Arrays

```go
array1 := [3]int{1,2,3}
array2 := array1
```

Arrays are **copied by value**.

Changing one array does **not** affect the other.

```
array1
↓

[1 2 3]

array2
↓

[1 2 3]
```

Independent copies.

---

# Slices

## What is a Slice?

A **slice** is a dynamic view over an underlying array.

Unlike arrays:

* Size can grow.
* Size can shrink.
* Uses an underlying array internally.
* Most Go programs use slices instead of arrays.

Declaration

```go
slice := []int{10,20,30}
```

---

## append()

Adds new elements.

```go
slice = append(slice, 40)
```

Output

```
[10 20 30 40]
```

Multiple values

```go
slice = append(slice, 50,60,70)
```

---

## Length

```go
len(slice)
```

Returns current number of elements.

---

## Capacity

```go
cap(slice)
```

Returns maximum elements the current underlying array can hold before Go allocates a new array.

Example

```
Length = 5

Capacity = 8
```

Capacity grows automatically when needed.

---

## Creating a Slice with make()

```go
values := make([]int,3,5)
```

Meaning

```
Length = 3

Capacity = 5
```

Initially

```
[0 0 0]
```

---

## Slice from an Array

```go
data := [5]int{10,20,30,40,50}

s := data[1:4]
```

Result

```
[20 30 40]
```

Rules

```
start included

end excluded
```

---

## Common Slice Expressions

```
slice[1:4]

slice[:3]

slice[2:]

slice[:]
```

Examples

```
Original

[10 20 30 40 50]

slice[1:4]

[20 30 40]

slice[:3]

[10 20 30]

slice[2:]

[30 40 50]

slice[:]

[10 20 30 40 50]
```

---

## Slice Assignment

```go
s1 := []int{1,2,3}

s2 := s1
```

Unlike arrays, slices share the same underlying array.

Changing

```go
s2[0] = 999
```

Also changes

```
s1
```

because both point to the same memory.

---

## copy()

Creates an independent copy.

```go
duplicate := make([]int,len(original))

copy(duplicate, original)
```

Now both slices are completely separate.

Changing one will not affect the other.

---

## Appending One Slice to Another

```go
a := []int{1,2}

b := []int{3,4}

a = append(a,b...)
```

Output

```
[1 2 3 4]
```

`...` expands the slice into individual elements.

---

## Delete an Element

Go has no built-in delete function for slices.

Use

```go
slice = append(slice[:index], slice[index+1:]...)
```

Example

```
Before

[10 20 30 40 50]

Delete index 2

After

[10 20 40 50]
```

---

## Insert an Element

```go
slice = append(slice[:index],
    append([]int{value}, slice[index:]...)...)
```

Example

```
Before

[10 20 40 50]

Insert 30

After

[10 20 30 40 50]
```

---

## Nil Slice vs Empty Slice

Nil slice

```go
var s []int
```

Empty slice

```go
s := []int{}
```

Both print

```
[]
```

Difference

```
nil slice == nil

true

empty slice == nil

false
```

---

## Capacity Growth

Appending elements increases capacity automatically.

Example

```
Append 1

len = 1

cap = 1

Append 2

len = 2

cap = 2

Append 3

len = 3

cap = 4

Append 5

len = 5

cap = 8
```

Go allocates a larger underlying array when the current capacity is exhausted.

---

# Arrays vs Slices

| Feature            | Array            | Slice                   |
| ------------------ | ---------------- | ----------------------- |
| Size               | Fixed            | Dynamic                 |
| Memory Allocation  | Immediate        | Uses underlying array   |
| Can Grow           | No               | Yes                     |
| Value or Reference | Value Type       | Reference-like Type     |
| Copy Behavior      | Independent Copy | Shares Underlying Array |
| Most Common        | Rare             | Very Common             |

---

# Methods and Built-in Functions Used

| Function   | Purpose                                            |
| ---------- | -------------------------------------------------- |
| `len()`    | Returns number of elements                         |
| `cap()`    | Returns slice capacity                             |
| `append()` | Adds elements to a slice                           |
| `copy()`   | Copies elements into another slice                 |
| `make()`   | Creates slices with predefined length and capacity |
| `range`    | Iterates over arrays and slices                    |

---

# Key Takeaways

* Arrays have a fixed size and are copied by value.
* Slices are dynamic and are built on top of arrays.
* `append()` is the primary way to grow a slice.
* `copy()` creates an independent slice.
* `len()` returns the current number of elements.
* `cap()` returns the allocated storage capacity.
* `make()` allocates memory for slices.
* Slices created through assignment share the same underlying array.
* Deletion and insertion are performed using combinations of slicing and `append()`.
* Slices are used far more frequently than arrays in production Go applications.
