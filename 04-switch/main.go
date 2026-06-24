package main

import "fmt"

func main() {

	fmt.Println("############# Switch Case Practice ##############")

	name := ""
	fmt.Print("Enter your name : ")
	fmt.Scan(&name)

	timeline := ""
	fmt.Print("Enter timeline : ")
	fmt.Scan(&timeline)

	switch timeline {
		case "Morning":
			fmt.Println("Good morning", name)
		case "Afternoon":
			fmt.Println("Good Afternoon", name)
		case "Evening":
			fmt.Println("Good evening", name)
		default:
			fmt.Println("Hello ", name)
	}
}

/*
=========================================================
fmt.Scan() vs fmt.Scanln()
=========================================================

1. fmt.Scan()

- Reads input token by token.
- Tokens are separated by:
	* Space
	* Tab
	* New Line (\n)

Example:

	var firstName string
	var lastName string

	fmt.Scan(&firstName, &lastName)

Input:
	Shaad Bangi

Output:
	firstName = Shaad
	lastName  = Bangi

---------------------------------------------------------

2. fmt.Scanln()

- Also reads input token by token.
- Stops reading when Enter (\n) is pressed.
- Typically used when all expected inputs should be on the same line.

Example:

	var firstName string
	var lastName string

	fmt.Scanln(&firstName, &lastName)

Input:
	Shaad Bangi

Output:
	firstName = Shaad
	lastName  = Bangi

---------------------------------------------------------

Important:

Both Scan() and Scanln() split input on spaces.

Example:

	var name string

	fmt.Scan(&name)

Input:
	Shaad Bangi

Output:
	name = Shaad

---------------------------------------------------------

Similarly:

	var name string

	fmt.Scanln(&name)

Input:
	Shaad Bangi

Output:
	name = Shaad

---------------------------------------------------------

Why did timeline get skipped?

Example:

	var name string
	var timeline string

	fmt.Scan(&name)

Input:
	Shaad Bangi

Result:
	name = Shaad

Remaining Input Buffer:
	Bangi

Next:

	fmt.Scan(&timeline)

timeline automatically gets:
	Bangi

So Go appears to skip asking for input.

---------------------------------------------------------

To read complete lines (full names, addresses, messages):

Use bufio.Reader

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

Input:
	Shaad Bangi

Output:
	name = "Shaad Bangi\n"

Remove newline:

	name = strings.TrimSpace(name)

---------------------------------------------------------

Rule of Thumb

fmt.Scan()
	- Numbers
	- Menu Choices
	- Single Words

fmt.Scanln()
	- Multiple values on one line

bufio.Reader + ReadString('\n')
	- Full Names
	- Addresses
	- Sentences
	- User Messages

=========================================================
*/