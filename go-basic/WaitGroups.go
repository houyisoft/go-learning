package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int) {
	fmt.Printf("worker2 %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("worker2 %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker2(i)
		}()
	}

	wg.Wait()

}
