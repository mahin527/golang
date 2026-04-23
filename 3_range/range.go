package main

import (
	"fmt"
)

// iterating over data structures
func main() {
	fmt.Println("Range")

	nums := []int{2, 3, 4, 5, 6}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// Range
	sum := 0
	// for _, num := range nums {
	for index, num := range nums {
		// fmt.Println(num)
		sum = sum + num
		fmt.Println(index, sum)
	}

	m4 := map[string]string{"name": "John", "lname": "Doe"}

	for k, v := range m4 {
		fmt.Println(k, v, ",")
	}
	// unicode code point rune
	// i -> starting byte of rune
	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}

}
