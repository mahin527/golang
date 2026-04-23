package main

import "fmt"

// func add(a int, b int) int {
func add(a, b int) int {
	return a + b
}

func getLang() (string, string, string) {
	return "Golang", "Javascript", "python"
}

func processIt(fn func(a int) int) {
	fn(1)
}

func main() {
	// function call
	result := add(10, 20)
	fmt.Println(result)

	// fmt.Println(getLang())

	lang1, lang2, _ := getLang()
	fmt.Println(lang1, lang2)
	// fmt.Println(lang3)

	fn := func(a int) int {
		return 2
	}

	processIt(fn)
}
