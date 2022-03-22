package main

import (
	"fmt"
	"time"
)

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
	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)
		go setInterval(i, Timer)
		cmd := <-Timer
		fmt.Println(cmd)
		if cmd == "Exit" {
			break
		}
	}
}
