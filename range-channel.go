package main

import "fmt"

func channelRange() {

	// buffered channel
	var ch = make(chan int, 10)

	go func() {
		for i := 0; i < cap(ch); i++ {
			ch <- i
		}
		// need to close otherwise it will block
		close(ch)
	}()

	for value := range ch {
		fmt.Println(value)
	}
}
