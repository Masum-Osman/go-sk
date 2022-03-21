package main

import (
	"fmt"
	"sync"
	"time"
)

func worker() {
	dt := time.Now()
	fmt.Println(dt.Format(time.UnixDate))
	time.Sleep(time.Second)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker()
		}()
	}

	wg.Wait()

}
