package main

import (
	"fmt"
	"time"
)

func timer(s int) {
	for {
		if s <= 0 {
			break
		} else {
			fmt.Println(s)
			time.Sleep(1 * time.Second)
			s--
		}
	}
}

func main() {
	go timer(10)
}
