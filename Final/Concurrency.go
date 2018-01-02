package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	Start := time.Now()

	p("Start :", Start)

	for i := 1; i <= 2000; i++ {
		if i%2 == 1 {
			go p(i)
		}
	}

	time.Sleep(100 * time.Millisecond)

	End := time.Now()
	p("Stop : ", End)
	t := time.Now()
	elapsed := t.Sub(Start)
	p("Waktu proses :", elapsed)
}
