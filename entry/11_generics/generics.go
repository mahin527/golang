package main

import "fmt"

func printSlice[T any](items []T) {
	// func printSlice[T interface{}](items []T) {
	// func printSlice[T string | int](items []T) {
	for index, item := range items {
		fmt.Println(index, ":", item)
	}
}

// LIFO (last in first out)
type stack struct {
	elements []int
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	names := []string{"golang", "java", "python"}

	printSlice(names)
	printSlice(nums)

	myStack := stack{
		elements: []int{5, 6, 7, 88, 90},
	}

	fmt.Println(myStack)
}
