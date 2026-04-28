/*
Interface is a contract: who will implement which method.

Interface → contract (Any struct can implement it)

*/

package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

// open close principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGtw := razorpay{}
	// stripePaymentGtw := stripe{}
	// razorpayPaymentGtw.pay(amount)
	// stripePaymentGtw.pay(amount)
	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment in razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment in stripe", amount)

}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose", amount)
}

func (f fakepayment) refund(amount float32, account string) {

}

func main() {
	// stripePaymentGtw := stripe{}
	fakeGtw := fakepayment{}

	newPayment := payment{
		// gateway: stripePaymentGtw,
		gateway: fakeGtw,
	}

	newPayment.makePayment(100.00)
}
