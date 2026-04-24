package main

import "fmt"

// by value
func changeNum(num int) {
	num = 5
	fmt.Println("In ChangeNum: ", num)
}

// by reference
func changeNum2(num2 *int) {
	*num2 = 5
	fmt.Println("In ChangeNum2: ", *num2)
}

func main() {

	num := 1
	changeNum(num)

	fmt.Println("In Main: ", num)
	fmt.Println("In Main Menory : ", &num)

	num2 := 1
	changeNum2(&num2)
	fmt.Println("In Main: ", num2)

}
