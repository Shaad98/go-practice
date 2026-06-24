package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var age uint
	fmt.Print("Enter your age : ")
	fmt.Scan(&age)

	fmt.Println(age)

	if age >= 18 {
		fmt.Println("You can apply for Driving License")
	} else {
		fmt.Println("You can't apply for Driving License")
	}

	fmt.Println("\n______________________________________\n")

	var percentage float32

	fmt.Print("Enter your percentage : ")
	fmt.Scan(&percentage)

	if percentage < 0 || percentage > 100 {
		fmt.Println("Invalid percentage")
	} else if percentage < 35 {
		fmt.Println("You have failed!")
	} else if percentage < 50 {
		fmt.Println("You need to improve yourself!")
	} else if percentage <= 75 {
		fmt.Println("You are an average student!")
	} else {
		fmt.Println("You got excellent marks!")
	}

	fmt.Println("\n______________________________________\n")

	// if with initialization
	if num := rand.Intn(100); num > 50 {
		fmt.Println("Greater than 50:", num)
	} else if num > 25 {
		fmt.Println("Greater than 25:", num)
	} else {
		fmt.Println("25 or below:", num)
	}

	// fmt.Println(num) // ❌ Compile Error: undefined: num
}

// Syntax:
//
// if initialization; condition {
//     // if block
// } else if condition {
//     // else if block
// } else {
//     // else block
// }
//
// --------------------------------------------------------
//
// Example:
//
// if num := rand.Intn(100); num > 50 {
//     fmt.Println(num)
// } else {
//     fmt.Println(num)
// }
//
// Compiler conceptually treats it like:
//
// {
//     num := rand.Intn(100) // Create variable
//
//     if num > 50 {
//         fmt.Println(num)
//     } else {
//         fmt.Println(num)
//     }
//
//     // num's scope ends here
// }
//
// fmt.Println(num) // ❌ Error: undefined: num
//
// --------------------------------------------------------
//
// Key Points:
//
// 1. The initialization executes only once.
// 2. The condition is checked after initialization.
// 3. The initialized variable is available in:
//      - if block
//      - else if block(s)
//      - else block
// 4. The variable is NOT accessible after the if statement.
// 5. This is mainly used for temporary variables that are
//    needed only inside the if statement.
//
// Common real-world example:
//
// if err := saveUser(user); err != nil {
//     fmt.Println(err)
//     return
// }
//
// err exists only inside this if statement.