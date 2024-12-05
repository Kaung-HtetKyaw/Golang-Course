package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func raceCondition() {
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Go Routine: ", runtime.NumGoroutine())
	var counter int64

	var numGoRoutines = 100

	var wg sync.WaitGroup
	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func() {

			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("Go Routines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Go Routines: ", runtime.NumGoroutine())

	fmt.Println("Counter: ", counter)
}
