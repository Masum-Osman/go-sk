package main

import (
	"fmt"
	"sync"
	"time"
)

func worker() {
	for {
		dt := time.Now()
		fmt.Println(dt.Format(time.UnixDate))
		time.Sleep(time.Second)
	}
}

func main1() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		worker()
	}()

	wg.Wait()
	wg.Done()

	s := 7
	for {
		fmt.Println(s)
		if s <= 0 {
			wg.Done()
			break
		} else {
			fmt.Println(s)
			time.Sleep(1 * time.Second)
			s--
		}
	}

}
