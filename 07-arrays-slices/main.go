package main

import "fmt"

func main() {

	fmt.Println("========== Arrays ==========")

	// Declare an array
	var numbers [5]int

	fmt.Println(numbers)

	// Initialize an array
	arr := [5]int{10, 20, 30, 40, 50}

	fmt.Println(arr)

	// Access element
	fmt.Println("First Element :", arr[0])

	// Update element
	arr[2] = 100

	fmt.Println(arr)

	fmt.Println()

	fmt.Println("Length of Array :", len(arr))

	fmt.Println()

	fmt.Println("Loop using for")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println()

	fmt.Println("Loop using range")

	for index, value := range arr {
		fmt.Printf("Index : %d Value : %d\n", index, value)
	}

	fmt.Println()

	fmt.Println("========== Array Copy ==========")

	array1 := [3]int{1, 2, 3}
	array2 := array1

	array2[0] = 100

	fmt.Println("Array1 :", array1)
	fmt.Println("Array2 :", array2)

	fmt.Println()

	fmt.Println("========== Slices ==========")

	slice := []int{10, 20, 30}

	fmt.Println(slice)

	fmt.Println()

	fmt.Println("Append")

	slice = append(slice, 40)
	slice = append(slice, 50)

	fmt.Println(slice)

	fmt.Println()

	fmt.Println("Length :", len(slice))
	fmt.Println("Capacity :", cap(slice))

	fmt.Println()

	fmt.Println("Slice from Array")

	data := [5]int{10, 20, 30, 40, 50}

	s := data[1:4]

	fmt.Println(s)

	fmt.Println()

	fmt.Println("Slice Copy (Shares Memory)")

	s1 := []int{1, 2, 3}

	s2 := s1

	s2[0] = 999

	fmt.Println("Slice1 :", s1)
	fmt.Println("Slice2 :", s2)

	fmt.Println()

	fmt.Println("Copy Slice")

	original := []int{10, 20, 30}

	duplicate := make([]int, len(original))

	copy(duplicate, original)

	duplicate[0] = 999

	fmt.Println(original)
	fmt.Println(duplicate)

	fmt.Println()

	fmt.Println("Make Slice")

	values := make([]int, 3, 5)

	fmt.Println(values)
	fmt.Println("Length :", len(values))
	fmt.Println("Capacity :", cap(values))

	fmt.Println()

	fmt.Println("Range Loop on Slice")

	for index, value := range slice {
		fmt.Printf("Index : %d Value : %d\n", index, value)
	}
}