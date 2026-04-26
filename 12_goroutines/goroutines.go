/*
A goroutine is a lightweight thread of execution in the Go runtime.
*/
package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// task(i)
		wg.Add(1)
		go task(i, &wg)

	}

	// time.Sleep(time.Second * 1)
	wg.Wait()
}
