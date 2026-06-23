# Go Notes

## Program Entry Point

Every executable Go program starts execution from the `main()` function in the `main` package.

```go
package main

func main() {
    // Program starts here
}
```

### Rules

* `package main` is required for an executable program.
* There must be **exactly one** `main()` function in a `main` package.
* A `main()` function in any other package is **not** the program's entry point.
* `main()` cannot be overridden or overloaded.
* `go run file.go` compiles only the specified file(s), so multiple files with `main()` are fine **if only one file is compiled**.
* `go run .` or `go build` compiles the entire package and fails if it contains more than one `main()` function (`main redeclared in this block`).

---

# Go Modules

A Go module is created using:

```bash
go mod init <module-name>
```

Example:

```bash
go mod init go-practice
```

This creates a `go.mod` file, which is similar to `package.json` in Node.js because it stores the module name and project dependencies.

---

## Why `go run .` Fails Without `go.mod`

If you run:

```bash
go run .
```

without a `go.mod` file, you'll see:

```text
go: cannot find main module
```

This happens because `go run .` means:

> Compile and run the **current package**, and Go expects that package to belong to a Go module.

Initialize the module first:

```bash
go mod init go-practice
go run .
```

---

## Running Go Programs

### 1. Run a Single File

```bash
go run main.go
```

Only `main.go` is compiled.

If `main.go` calls a function from another file, for example:

```go
Greet()
```

and `Greet()` is defined in `helper.go`, you'll get:

```text
undefined: Greet
```

because `helper.go` was **not** compiled.

---

### 2. Run Multiple Files (Without a Module)

You can compile multiple files together:

```bash
go run main.go helper.go
```

Go compiles only the files you specify.

---

### 3. Run an Entire Project (Recommended)

After creating a module:

```bash
go mod init go-practice
```

Run the whole package:

```bash
go run .
```

Go automatically compiles **all `.go` files** in the current package, so functions from other files (like `Greet()`) are available without specifying every file.

---

## Summary

| Command                    | What Gets Compiled                                         |
| -------------------------- | ---------------------------------------------------------- |
| `go run main.go`           | Only `main.go`                                             |
| `go run main.go helper.go` | Only the specified files                                   |
| `go run .`                 | All `.go` files in the current package (requires `go.mod`) |

### Recommendation

* Small practice programs → `go run main.go`
* Multiple files → `go run main.go helper.go`
* Real projects → `go mod init` then `go run .` or `go build`
