/*
In Go, a struct is a data container.(struct -> Data Structure)

Struct → data container (like object, no inheritance).

Go doesn't have inheritance, but it does have composition.

Inheritance means: a class/struct gets properties + behavior from another class.

Composition means: a struct/class uses another struct inside it.
There is no inheritance in Go, so composition is used.

*/

package main

import (
	"fmt"
	"time"
)

// customer
type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	// struct embed
	customer
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// user
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	fmt.Println("Structs")

	order1 := order{
		id:     "1",
		amount: 21.09,
		status: "panding",
	}

	// if you don't set any fild , default value is zero
	// int -> 0, float-> 0, string -> "", bool -> false

	order1.createdAt = time.Now()

	order2 := order{
		id:        "2",
		amount:    28.78,
		status:    "delivered",
		createdAt: time.Now(),
	}

	fmt.Println(order1)
	fmt.Println("order1.status: ", order1.status)
	order1.changeStatus("confirmed")
	fmt.Println("order1.status: ", order1.status)

	fmt.Println(order2)
	order1.status = "received"
	fmt.Println(order2.status)
	fmt.Println(order1)

	// user
	p := Person{
		Name: "Mahin",
		Age:  24,
	}
	p.Greet() // Output: Hello, my name is Mahin

	myOrder := newOrder("1", 50.32, "panding")

	fmt.Println(myOrder)
	fmt.Println(myOrder.amount)

	language := struct {
		name   string
		isGood bool
	}{
		"Golang",
		true,
	}

	fmt.Println(language)

	customer1 := customer{
		name:  "john doe",
		phone: "01324678890",
	}
	order3 := order{
		id:        "3",
		amount:    23.76,
		status:    "Delivered",
		createdAt: time.Now(),
		customer:  customer1,
	}

	fmt.Println(order3)

	order4 := order{
		id:        "4",
		amount:    63.06,
		status:    "received",
		createdAt: time.Now(),
		customer: customer{
			name:  "Mahin",
			phone: "098487556",
		},
	}

	fmt.Println(order4)
}
