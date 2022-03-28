package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("Starting")
}
func TimeMachine() chan string {
	ch := make(chan string)
	go func() {
		for {
			fmt.Println(time.Now())
			select {
			case cmd := <-ch:
				fmt.Println(time.Now())
				fmt.Println(cmd)
				wg.Done()
				return
				// break
			case <-time.After(1 * time.Second):
				break
			}
		}
	}()
	return ch
}
func main() {
	wg.Add(1)
	ch := TimeMachine()
	time.Sleep(20 * time.Second)
	ch <- "Done"
	wg.Wait()
}
