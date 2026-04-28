package main

import "fmt"

// enums -> enumerated types

// type OrdreStatus int

// const (
// 	Received OrdreStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrdreStatus string

const (
	Received  OrdreStatus = "received"
	Confirmed OrdreStatus = "confirmed"
	Prepared  OrdreStatus = "prepared"
	Delivered OrdreStatus = "delivered"
)

// func changingOrderStatus(status string) {
// 	fmt.Println("changing order status to:", status)
// }

func changingOrderStatus(status OrdreStatus) {
	fmt.Println("changing order status to:", status)
}

func main() {
	// changingOrderStatus("received")

	changingOrderStatus(Prepared)

}
