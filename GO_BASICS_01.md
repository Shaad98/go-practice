# Go Basics - Notes

## 1. Do I need to initialize every Go project?

**Recommended:** Yes.

Whenever you start a new Go project, initialize it as a Go module.

```bash
mkdir go-practice
cd go-practice
go mod init github.com/your-username/go-practice
```

If you don't have a GitHub repository yet, you can simply use:

```bash
go mod init go-practice
```

For very small practice programs, you can write a single `.go` file and run it without `go mod init`, but for real projects always use Go modules.

---

# 2. What is `go.mod`?

After running:

```bash
go mod init go-practice
```

Go creates:

```go
module go-practice

go 1.24
```

`go.mod` contains:

* Module (project) name
* Go version
* Project dependencies

It is similar to **package.json** in Node.js, but much simpler.

---

# 3. What is `go.sum`?

When you add external libraries, Go creates a `go.sum` file.

Purpose:

* Stores checksums of downloaded dependencies.
* Verifies dependency integrity.
* Should be committed to Git.

---

# 4. Which package can be executed?

Only a package named **main** with a **main()** function can be executed.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```

Run:

```bash
go run main.go
```

---

# 5. Why did I get this error?

```
package command-line-arguments is not a main package
```

Reason:

* The file is not using `package main`, or
* It does not contain a `main()` function.

Only `package main` can be executed.

---

# 6. Can I create multiple packages in the same folder?

**No.**

One folder can contain only one package.

❌ Wrong

```
project/
├── main.go      // package main
└── helper.go    // package helper
```

This causes an error.

---

✅ Correct

```
project/
├── main.go

helper/
└── helper.go
```

`main.go`

```go
package main
```

`helper.go`

```go
package helper
```

---

# 7. Rule to remember

```
One Folder = One Package
```

Every `.go` file inside the same folder must have the same package name.

---

# 8. Should folder name and package name be the same?

They **do not have to be the same**, but they **should be the same**.

Recommended:

```
helper/
    helper.go
```

```go
package helper
```

This is the Go convention and is used in most projects.

---

Allowed but not recommended:

```
helper/
    helper.go
```

```go
package utils
```

This works, but it can make the code confusing.

---

# 9. Recommended project structure

```
go-practice/
│
├── go.mod
├── go.sum
├── main.go
│
├── handlers/
├── services/
├── models/
├── repository/
├── utils/
└── config/
```

Each folder represents one package.

---

# 10. Common Go commands

Initialize a project:

```bash
go mod init go-practice
```

Run the project:

```bash
go run .
```

Run a single file:

```bash
go run main.go
```

Download and clean dependencies:

```bash
go mod tidy
```

Build an executable:

```bash
go build
```

---

# Quick Summary

* Initialize every real project using `go mod init`.
* `go.mod` manages the project and dependencies.
* `go.sum` stores dependency checksums.
* Only `package main` with `main()` can be executed.
* One folder can contain only one package.
* Folder name and package name should be the same.
* Follow Go conventions to keep your code clean and maintainable.
