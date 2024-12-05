package main

import "fmt"

func ninja10Seven() {
	var channel = make(chan int)

	for x := 0; x < 10; x++ {
		go func() {
			for i := 0; i < 10; i++ {
				channel <- i
			}

		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-channel)
	}

}
