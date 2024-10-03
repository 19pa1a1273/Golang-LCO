package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channels in golang - LearnCodeOnline.in")

	myCh := make(chan int, 2) // Creates a buffered channel with a capacity of 2

	// {From ChatGPT}
	// ch <- 1
	// ch <- 2  // These sends won't block because the channel has space

	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)

	// Outside of the Box means Recieve
	// R ONLY

	// Eg:
	// return <-ch  // Only receiving is allowed

	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen) //true means sent the value, False means close 0 is sent
		fmt.Println(val)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// Inside the Box means Sending
	// SEND ONLY

	// Eg:
	// ch <- 5  // Only sending is allowed

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}

// Buffered Channels:

// 1. Channels can be buffered,
// 	  which means they can hold a specified number of values before any send operation blocks.
// 2. If a channel is not buffered, sends block until the data is received.

// close(ch):

// 1. Closing a channel indicates that no more values will be sent on it,
// 	  and this can be useful for signaling completion.
// 2. After closing, any attempt to send on the channel will cause a panic,
// 	  but receiving from a closed channel continues until the buffer is drained.

// Advantages of Channels

// 1. Synchronization:
//    	Channels can be used to synchronize the execution of goroutines by
// 		ensuring that certain operations happen in a particular order.
// 2. Communication:
// 		Channels provide a safe and efficient way to communicate between goroutines,
// 		avoiding the need for explicit locking.
