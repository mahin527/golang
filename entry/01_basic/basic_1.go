package main

import (
	"fmt"
	"slices"
	"time"
)

func main() {
	// print hello
	fmt.Println("Hello, Developer!")

	// simple values
	// data type is : numbers (integers, floats, complex), strings, and booleans
	fmt.Println(5 + 4)
	fmt.Println(5.78 / 4.65)

	// variables
	var name string = "Mahin"
	var age int = 23
	var isActive bool = true

	fmt.Println(name, age)
	fmt.Println(isActive)

	var name5 = "Hasan"

	// shorthand syntex of declaring variables

	name2 := "Mahin"
	fmt.Println(name2, name5)

	var name3 string

	name3 = "golang"
	fmt.Println(name3)

	// constants

	var studentName1 = "Akij Hasan"
	studentName1 = "Raihan Ahmad"
	fmt.Println(studentName1)

	const studentName2 = "Akij Hasan"
	// studentName2 = "Raihan Ahmad" // Error because it is constant var
	// cannot assign to studentName2 (neither addressable nor a map index expression)compilerUnassignableOperand
	fmt.Println(studentName2)

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(port, host)

	// for -> only construct in go 'for' looping
	// while loop implement
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop

	// for {
	// 	println("1")
	// }

	// classic for loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		break
		// continue
	}

	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Persom is an adult")
	// } else {
	// 	fmt.Println("Persom is not an adult")
	// }

	role := "admin"
	hasPermission := true

	// ==, ||, &&
	if role == "admin" && hasPermission {
		fmt.Println("You have access for this application")
	}

	// go does not have 'ternary operators', will have to use normal 'if else'

	// switch
	var yourAge int = 20

	switch yourAge {
	case 18:
		fmt.Println("You are an adult")

	case 17:
		fmt.Println("You are not and adult")
	}

	var day int = 3

	switch day {
	case 1:
		fmt.Println("Monday")
		fallthrough // Explicitly continues to the next case
	case 2:
		fmt.Println("Tuesday") // Prints even if day was 1

	default:
		fmt.Println("Other Day!")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("it's weekend!")
	default:
		fmt.Println("it's workday!")
	}

	// type switch
	WhoAmI := func(i interface{}) {
		// WhoAmI := func(i any{}) {
		switch t := i.(type) {
		// switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a bool")
		default:
			fmt.Println("Other", t)
			fmt.Println("Other")

		}
	}

	WhoAmI(40.09)
	WhoAmI(42)
	WhoAmI("Mahin")
	WhoAmI(true)

	//======================
	//======== Array =======
	//======================

	// numbered sequence of specific length
	var names [4]string

	// length of nums
	fmt.Println(len(names))

	names[0] = "mahin"

	fmt.Println(names[0])

	fmt.Println(names)

	// to declare it in single line
	nums := [4]int{1, 2, 3, 4}
	fmt.Println(nums)

	// 2d arrays
	myNum := [2][2]int{{1, 2}, {8, 6}}
	fmt.Println(myNum)

	/*
		you can use array if :

		1. fixed size, that is predictable
		2. memory optimization
		3. constant time access
	*/

	// slice
	// slice -> dynamic
	// most used constructor in 'go'
	// + useful methods

	// uninitialized slice is nil
	var mySlice []int

	fmt.Println(mySlice)

	fmt.Println(mySlice != nil)
	fmt.Println(len(mySlice))

	// if you want value != nil
	var mySlice2 = make([]int, 3, 5)
	fmt.Println(mySlice2)

	fmt.Println(mySlice2 != nil)

	// capacity -> maximum numbers of elements can fit
	fmt.Println(cap(mySlice2))

	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 3)
	mySlice2 = append(mySlice2, 4)
	mySlice2 = append(mySlice2, 5)
	fmt.Println(cap(mySlice2))
	fmt.Println(mySlice2)

	// copy function
	fruits := []string{"Apple", "Banana", "Litchi", "Mango", "Orange", "PineApple"}
	fruits2 := make([]string, len(fruits))
	copy(fruits2, fruits)

	fmt.Println(fruits)  // original
	fmt.Println(fruits2) // full copy

	// slice operator
	// fmt.Println(fruits[0:4])
	fmt.Println(fruits[:4])
	fmt.Println("test")
	fmt.Println(fruits[:])

	fmt.Println(slices.Equal(fruits, fruits2))

	// 2d slices
	fmt.Println("2d slices")

	var nums3 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums3)
}
