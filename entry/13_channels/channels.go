/*
Go's channels are special pipelines for exchanging data between goroutines.
It is the most powerful feature for handling concurrency.
*/

package main

import (
	"fmt"
	// "time"
)

// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("processing number:", num)
// 		time.Sleep(time.Second)
// 	}

// 	// fmt.Println("processing number:", <-numChan)
// }

// func sum(result chan int, num1 int, num2 int) {
// 	sumResult := num1 + num2
// 	result <- sumResult
// }

// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Processing...")
// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Ping"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}

	}

	// emailChan := make(chan string, 20)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// go task(done)
	// <-done

	// emailChan <- "1@gmail.com"
	// emailChan <- "2@proton.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	// fmt.Println("Done sending!")

	// this is important
	// close(emailChan)
	// <-done

	// result := make(chan int)

	// go sum(result, 5, 6)

	// recieveResult := <-result

	// fmt.Println(recieveResult)
	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.IntN(100)
	// }

	// numChan <- 5

}
