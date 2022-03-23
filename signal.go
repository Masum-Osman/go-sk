package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker1() {
	dt := time.Now()
	fmt.Println(dt.Format(time.UnixDate))
	time.Sleep(time.Second)
}

func main10() {
	var wg sync.WaitGroup

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	// go func() {

	// 	sig := <-sigs
	// 	fmt.Println()
	// 	fmt.Println(sig)
	// 	done <- true
	// }()

	go func() {
		defer wg.Done()
		worker1()
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
