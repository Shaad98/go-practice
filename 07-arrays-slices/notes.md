# Arrays and Slices in Go

Arrays and slices are used to store collections of values. While they appear similar, they behave very differently.

Arrays have a fixed size, whereas slices are dynamic and can grow as needed.

---

# Arrays

An array is a collection of elements of the same type with a fixed length.

## Syntax

```go
var arr [5]int
```

or

```go
arr := [5]int{1,2,3,4,5}
```

---

## Characteristics

- Fixed size
- Stores elements of the same type
- Value type
- Cannot change size after creation

Example

```go
numbers := [3]int{10,20,30}
```

---

## Access Elements

```go
fmt.Println(numbers[0])
```

Output

```
10
```

---

## Modify Elements

```go
numbers[1] = 100
```

---

## Length

```go
len(numbers)
```

---

## Looping

Using for

```go
for i:=0;i<len(numbers);i++{
    fmt.Println(numbers[i])
}
```

Using range

```go
for index,value := range numbers{
    fmt.Println(index,value)
}
```

---

# Arrays are Value Types

Assigning one array to another copies every element.

```go
a := [3]int{1,2,3}

b := a

b[0] = 100
```

Result

```
a = [1 2 3]

b = [100 2 3]
```

The original array is unchanged.

---

# Slices

A slice is a flexible view over an array.

Unlike arrays, slices can grow and shrink.

Syntax

```go
numbers := []int{10,20,30}
```

---

# Append

```go
numbers = append(numbers,40)
```

Result

```
[10 20 30 40]
```

---

# Length

```go
len(slice)
```

Returns the number of elements.

---

# Capacity

```go
cap(slice)
```

Returns how many elements can fit before Go allocates a new backing array.

---

# Creating Slice from Array

```go
arr := [5]int{10,20,30,40,50}

slice := arr[1:4]
```

Result

```
[20 30 40]
```

---

# Slices Share Memory

```go
s1 := []int{1,2,3}

s2 := s1

s2[0] = 100
```

Result

```
s1 = [100 2 3]

s2 = [100 2 3]
```

Both slices reference the same backing array.

---

# Copy Slice

To create an independent copy

```go
destination := make([]int,len(source))

copy(destination,source)
```

Now modifying one slice does not affect the other.

---

# make()

The built-in `make()` function creates slices.

```go
slice := make([]int,3,5)
```

Meaning

```
Length = 3

Capacity = 5
```

---

# Difference Between Array and Slice

| Arrays | Slices |
|---------|---------|
| Fixed size | Dynamic size |
| Value type | Reference-like type |
| Copies entire array | Shares backing array |
| Rarely used directly | Used everywhere in Go |
| Size is part of type | Size is not part of type |

---

# Array Memory

```
Array

+----+----+----+----+
|10  |20  |30  |40  |
+----+----+----+----+
```

Copying

```
Array A

+----+----+----+

↓

Array B

+----+----+----+
```

Each array has its own memory.

---

# Slice Memory

```
Array

+----+----+----+----+----+
|10  |20  |30  |40  |50  |
+----+----+----+----+----+
      ↑
      │
Slice
```

The slice points to a portion of the array instead of copying it.

---

# Built-in Functions

| Function | Purpose |
|----------|---------|
| len() | Number of elements |
| cap() | Capacity |
| append() | Add elements |
| copy() | Copy slice |
| make() | Create slice |

---

# Interview Questions

## Are arrays and slices the same?

No.

Arrays have fixed size.

Slices are dynamic.

---

## Which one is used more?

Slices.

Almost every Go API uses slices instead of arrays.

---

## Why do slices affect each other?

Because multiple slices may share the same backing array.

---

## When should copy() be used?

When you need an independent slice that does not share memory with the original.

---

## Can an array grow?

No.

Arrays always have a fixed length.

---

# Key Takeaways

- Arrays have a fixed size.
- Arrays are value types.
- Slices are built on top of arrays.
- Slices are dynamic.
- append() adds elements.
- make() creates slices.
- copy() creates independent slices.
- len() returns element count.
- cap() returns total capacity.
- Slices share memory unless explicitly copied.