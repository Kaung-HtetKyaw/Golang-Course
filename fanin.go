package main

import (
	"fmt"
)

func producer(ch chan<- int, d int) {
	for i := 0; i < 3; i++ {
		ch <- i + d
	}
	close(ch)
}

func fanIn(input1, input2 <-chan int, output chan<- int) {
	for {
		select {
		case i1, ok1 := <-input1:
			if ok1 {
				output <- i1
			}
		case i2, ok2 := <-input2:
			if ok2 {
				output <- i2
			}
		}
	}
}

func goFanIn() {
	output := make(chan int)
	input1 := make(chan int)
	input2 := make(chan int)

	go producer(input1, 0)
	go producer(input2, 10)
	go fanIn(input1, input2, output)

	for n := range output {
		fmt.Println(n)
	}

	close(output)
}
