package main

import (
	"fmt"
	"time"
)

func TimeMachine() (chan string, chan string) {
	ch := make(chan string)
	done := make(chan string)
	//var cmd string
	go func() {
		for {
			// 			fmt.Println(time.Now())
			// 			time.Sleep(10 * time.Second)
			select {
			case cmd := <-ch:
				fmt.Println(cmd)
				done <- "Done"
				return
			case <-time.After(1 * time.Second):
				fmt.Println(time.Now())
			}
		}

	}()
	return ch, done
}

func main() {
	ch, done := TimeMachine()
	//time.Sleep(1000 * 10)
	// fmt.Println(ch)
	// 	fmt.Println("Exit")
	time.Sleep(10 * time.Second)
	ch <- "Exit"
	fmt.Println(<-done)
}
