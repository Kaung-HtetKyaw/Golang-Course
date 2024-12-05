package main

import "fmt"

func ninja10Six() {
	var channel = make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel)
	}()

	for value := range channel {
		fmt.Println(value)
	}

}
