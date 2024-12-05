package main

import "fmt"

func channelSelect() {
	var even = make(chan int)
	var odd = make(chan int)
	var quit = make(chan int)

	go sendChannelSelect(even, odd, quit)

	receiveChannelSelect(even, odd, quit)

}

func receiveChannelSelect(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even: ", v)
		case v := <-odd:
			fmt.Println("Odd: ", v)
		case v := <-quit:
			fmt.Println("Quit: ", v)
			return
		}
	}
}

func sendChannelSelect(even, odd, quit chan<- int) {
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
	close(even)
	close(odd)
	close(quit)
}
