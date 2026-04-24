package main

import (
	"fmt"
	"time"
)

// order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func main() {
	fmt.Println("Structs")

	order1 := order{
		id:     "1",
		amount: 21.09,
		status: "panding",
	}

	order1.createdAt = time.Now()

	fmt.Println(order1)
}
