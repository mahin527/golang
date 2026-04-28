package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Map")

	// Map

	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "Golang"
	m["area"] = "Backend"
	m["package"] = "Main"

	// get an element
	fmt.Println(m)
	fmt.Println(m["name"], m["area"])

	// IMP: if key does not exists in the map then it returns zero
	fmt.Println(m["phone"])

	m2 := make(map[string]int)

	m2["age"] = 20
	delete(m, "package")
	fmt.Println(m)

	fmt.Println(m2["age"])
	fmt.Println(m2["num"])

	// length
	fmt.Println(len(m))
	fmt.Println(len(m2))

	clear(m)
	fmt.Println(m)

	// creating map without map
	m3 := map[string]int{"price": 120, "order": 6}

	fmt.Println(m3)

	// _, ok := m3["price"]
	v, ok := m3["order"]
	fmt.Println("test")
	fmt.Println(v)
	if ok {
		fmt.Println("All ok!")
	} else {
		fmt.Println("something not ok!")
	}

	fmt.Println(maps.Equal(m2, m3))
}
