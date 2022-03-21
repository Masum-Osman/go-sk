package main

import (
	"fmt"
	"time"
)

func foo() {
	i := 0
	for true {
		dt := time.Now()
		i++
		fmt.Println("new goroutine: i = ", dt.Format(time.UnixDate))
		time.Sleep(time.Second)
	}
}

func main() {
	//Create a goroutine and start another task
	go foo()

	i := 0
	for true {
		i++
		fmt.Println("main goroutine: i = ", i)
		time.Sleep(time.Second)
	}
}
