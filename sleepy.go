package main

import (
	"fmt"
	"sync"
	"time"
)

/*
func setInterval(milliseconds int, timer chan string) chan string {
	if milliseconds > 20 {
		timer <- "Exit"
		return timer
	}
	current := time.Now()
	timer <- current.Format("2006-01-02 15:04:05")
	return timer
}

func main() {
	Timer := make(chan string)
	a := 10;
	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)
		go setInterval(i, Timer, a)
		cmd := <-Timer
		fmt.Println(cmd)
		if cmd == "Exit" {
			break
		}
	}
}


*/
var wg sync.WaitGroup

func setInterval(timer chan string) {
	defer wg.Done()
	fmt.Println("WHAT?")

	for {
		fmt.Println("LOOP?")

		cmd := <-timer
		fmt.Println(cmd)

		if cmd == "Exit" {
			fmt.Println("INSIDE IF?")
			// defer wg.Done()
			// break
		} else {
			fmt.Println("Not Exit?")
			current := time.Now()
			fmt.Println(current.String())
		}
	}
}

func routine() {
	defer wg.Done() // 3
	fmt.Println("routine finished")
}

func main() {
	Timer := make(chan string)

	wg.Add(1)

	go setInterval(Timer)
	time.Sleep(1 * time.Second)
	Timer <- "Not"
	time.Sleep(3 * time.Second)
	Timer <- "Exit"
	wg.Wait()

	/*
		for i := 1; i < 21; i++ {
			go setInterval(Timer)
			time.Sleep(1 * time.Second)
			if i == 20 {
				Timer <- "Exit"
			}
			wg.Wait()
		}


			wg.Add(1)    // 2
			go routine() // *
			wg.Wait()    // 4
			fmt.Println("main finished")
	*/
}
