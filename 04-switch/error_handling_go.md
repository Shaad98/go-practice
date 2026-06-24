# Error Handling in Go

## Java Approach

```java
try {
    String name = scanner.nextLine();
} catch(Exception e) {
    System.out.println(e);
}
```

## Go Approach

```go
name, err := reader.ReadString('\n')

if err != nil {
    fmt.Println(err)
    return
}
```

---

## Does Go Have Try-Catch?

Go does **not** have:

* ❌ `try`
* ❌ `catch`
* ❌ `throws`
* ❌ `finally`

Instead, Go handles errors by returning them as values.

---

## Error Handling Pattern

```go
result, err := someFunction()

if err != nil {
    // Handle error
    return
}
```

This is the standard way to handle errors in Go.

---

## Functions Often Return

Many Go functions return two values:

```go
(value, error)
```

Examples:

```go
file, err := os.Open("data.txt")
```

```go
user, err := getUser()
```

```go
data, err := json.Marshal(user)
```

The first value is the actual result, and the second value is an error (if one occurred).

---

## Panic and Recover

### panic()

```go
panic("Something went wrong")
```

* Stops normal program execution.
* Causes the application to crash if not recovered.

### recover()

```go
recover()
```

* Used to recover from a panic.
* Typically used inside a deferred function.

Example:

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

---

## Java vs Go

| Java      | Go                  |
| --------- | ------------------- |
| Exception | panic               |
| catch     | recover             |
| try-catch | error return values |

However, panic/recover is **not** used for normal error handling in Go.

---

## Recommended Go Style

```go
value, err := function()

if err != nil {
    // Handle error
}
```

This makes error handling explicit and easy to understand.

---

## Most Common Pattern in Go

```go
if err != nil {
    return err
}
```

You will see this pattern everywhere in production Go code.

---

## Key Takeaway

Go prefers:

```go
result, err := someFunction()

if err != nil {
    return err
}
```

instead of exception-based error handling.

Errors are returned as values and handled explicitly by the developer.
