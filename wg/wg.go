package main

import (
	"fmt"
	"sync"
	"time"
)

func goro() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			goro()

		}()
	}

	wg.Wait()
}
