# Go Pointers

Pointers are one of Go's most important concepts. A pointer stores the **memory address** of another variable instead of storing the actual value.

Unlike C/C++, Go supports pointers but **does not allow pointer arithmetic**, making programs safer.

---

# Why Use Pointers?

Pointers are mainly used to:

- Avoid copying large values.
- Modify variables inside functions.
- Share the same data between multiple functions.
- Improve performance when working with large structs.

---

# Pointer Syntax

```go
var p *int
```

- `*int` means **pointer to an integer**.
- Initially, a pointer contains `nil`.

---

# Address Operator (`&`)

The `&` operator returns the memory address of a variable.

Example:

```go
x := 10

fmt.Println(&x)
```

Example Output

```
0xc000012090
```

---

# Dereference Operator (`*`)

The `*` operator accesses the value stored at the memory address.

Example

```go
x := 10

p := &x

fmt.Println(*p)
```

Output

```
10
```

---

# Memory Representation

```
Variable x

+------+
|  10  |
+------+

Address
0xc000012090


Pointer p

+------------------+
|0xc000012090      |
+------------------+
```

`p` stores the address.

`*p` accesses the value stored at that address.

---

# Modifying Value Through Pointer

```go
x := 10

p := &x

*p = 50
```

Result

```
x = 50
```

Changing `*p` changes the original variable because both refer to the same memory location.

---

# Call by Value

Go passes function arguments **by value**.

Example

```go
func change(num int){
    num = 100
}
```

Calling

```go
x := 10

change(x)
```

Result

```
x = 10
```

The function receives a copy of `x`.

---

# Call by Reference (Using Pointers)

```go
func change(num *int){
    *num = 100
}
```

Calling

```go
change(&x)
```

Result

```
x = 100
```

The function receives the memory address of `x`, so it modifies the original variable.

---

# Nil Pointer

An uninitialized pointer has the value `nil`.

```go
var p *int

fmt.Println(p)
```

Output

```
<nil>
```

Trying to dereference a nil pointer causes a runtime panic.

```go
*p = 10
```

Result

```
panic: invalid memory address
```

---

# Allocating Memory with `new()`

The built-in `new()` function allocates memory and returns a pointer.

```go
num := new(int)
```

Initially

```
*num == 0
```

After

```go
*num = 500
```

```
*num == 500
```

---

# Difference Between Value and Pointer

| Call by Value | Call by Pointer |
|---------------|-----------------|
| Receives a copy | Receives memory address |
| Original value remains unchanged | Original value changes |
| Less efficient for large data | More efficient for large data |

---

# Important Operators

| Operator | Meaning |
|----------|---------|
| `&` | Get memory address |
| `*` | Dereference pointer |
| `new()` | Allocate memory and return pointer |

---

# Pointer Flow

```
x = 10

        &
        │
        ▼
+-------------+
| Address     |
+-------------+
        ▲
        │
        p

*p

↓

10
```

---

# Interview Questions

### What is a pointer?

A pointer is a variable that stores the memory address of another variable.

---

### What does `&` do?

Returns the address of a variable.

---

### What does `*` do?

Accesses the value stored at the address held by a pointer.

---

### Does Go support pointer arithmetic?

No.

Unlike C/C++, Go does not allow operations such as:

```go
p++
```

This restriction improves memory safety.

---

### Is Go truly call by reference?

No.

Go is always **pass-by-value**.

When you pass a pointer to a function, the pointer itself is copied, but both the original and copied pointers refer to the same memory location. This allows the function to modify the underlying value.

---

# Key Takeaways

- A pointer stores the memory address of another variable.
- `&` returns the address of a variable.
- `*` accesses the value stored at an address.
- Passing a pointer allows functions to modify the original value.
- Uninitialized pointers are `nil`.
- `new()` allocates memory and returns a pointer.
- Go does not support pointer arithmetic.
- Go is always pass-by-value, even when passing pointers.