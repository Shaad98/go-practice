# Reading Full Line Input with bufio.Reader

## Packages Used

### fmt

Used for console input and output.

```go
import "fmt"
```

Examples:

```go
fmt.Print("Enter your name: ")
fmt.Println("Hello", name)
```

---

### os

Provides access to operating system resources.

```go
import "os"
```

Example:

```go
os.Stdin
```

`os.Stdin` represents standard input (keyboard input).

---

### bufio

Provides buffered input/output operations.

```go
import "bufio"
```

Example:

```go
reader := bufio.NewReader(os.Stdin)
```

Used to read full lines of input.

---

## Example Code

```go
reader := bufio.NewReader(os.Stdin)

fmt.Print("Enter your full name: ")

name, err := reader.ReadString('\n')

if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println("Hello", name)
```

---

## Why Not Use fmt.Scan()?

`fmt.Scan()` reads input only until the first whitespace.

Input:

```text
Shaad Bangi
```

Using:

```go
fmt.Scan(&name)
```

Result:

```text
name = "Shaad"
```

Only the first word is stored.

---

## Using bufio.Reader

```go
reader := bufio.NewReader(os.Stdin)
```

Creates a Reader object that reads data from standard input (`os.Stdin`).

Java Equivalent:

```java
Scanner scanner = new Scanner(System.in);
```

---

## ReadString('\n')

```go
name, err := reader.ReadString('\n')
```

Reads characters until a newline character (`\n`) is encountered.

Input:

```text
Shaad Bangi
```

Stored value:

```text
Shaad Bangi\n
```

The newline character is included in the returned string.

---

## Multiple Return Values

`ReadString()` returns:

```go
(string, error)
```

Example:

```go
name, err := reader.ReadString('\n')
```

Possible values:

```go
name = "Shaad Bangi\n"
err = nil
```

or

```go
name = ""
err = someError
```

---

## Error Handling in Go

Java:

```java
try {
    String name = scanner.nextLine();
} catch(Exception e) {
    System.out.println(e);
}
```

Go:

```go
name, err := reader.ReadString('\n')

if err != nil {
    fmt.Println(err)
    return
}
```

---

## Does Go Have try-catch?

No.

Go does not provide:

* ❌ try
* ❌ catch
* ❌ throws
* ❌ finally

Instead, errors are returned as values.

---

## Common Error Handling Pattern

```go
result, err := someFunction()

if err != nil {
    // handle error
    return
}
```

Most common production pattern:

```go
if err != nil {
    return err
}
```

---

## Key Takeaways

* `fmt` → console input/output.
* `os` → access to operating system resources (`os.Stdin`).
* `bufio` → buffered input/output and full-line reading.
* `fmt.Scan()` reads only one word.
* `ReadString('\n')` reads the entire line.
* Go handles errors using returned error values.
* Normal Go programs use `if err != nil` instead of try-catch.
