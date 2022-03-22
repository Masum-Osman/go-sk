package main

import (
	"fmt"
	"time"
)

func foo() {
	for true {
		dt := time.Now()
		fmt.Println(dt.Format(time.UnixDate))
		time.Sleep(time.Second)
	}
}

func main0() {
	go foo()
	// foo()
	for true {
		// time.Sleep(time.Second)
		fmt.Println("N")
	}
}
