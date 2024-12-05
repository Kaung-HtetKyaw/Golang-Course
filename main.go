package main

import (
	"fmt"

	"example.com/name/package/dog"
)

func channelWontRun() {
	ch := make(chan int)

	ch <- 42

	fmt.Println(<-ch)
}

func channelWillRun() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}

func bufferChannelSuccessful() {
	ch := make(chan int, 2)

	ch <- 42

	fmt.Println(<-ch)
}

func main() {
	dog.Years()
}
