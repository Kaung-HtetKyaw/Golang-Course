package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func goContext() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error check 1: ", ctx.Err())
	fmt.Println("Number of Go Routines: ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working on job: ", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2: ", ctx.Err())
	fmt.Println("About to cancel")
	cancel()
	fmt.Println("Cancelled Context")

	fmt.Println("Error check 3: ", ctx.Err())
	fmt.Println("Num of Go Routines: ", runtime.NumGoroutine())
}
