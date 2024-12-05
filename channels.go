package main

import "fmt"

func goChannels() {
	var ch = make(chan int)

	go send(ch)

	receive(ch)

}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
