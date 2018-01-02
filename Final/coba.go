package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	tStart := time.Now()

	// time start
	p(tStart)

	// printing numbers
	for i := 0; i <= 100000; i++ {
		if i%2 == 1 {
			p(i)
		}
	}

	// sleep
	time.Sleep(100 * time.Millisecond)

	tEnd := time.Now()
	// time stop
	p(tEnd)

	// calculate elapsed time
	tElapsed := tEnd.Sub(tStart)
	p(tElapsed)
}
